package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Astemirdum/lavka/pkg/functools"

	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Astemirdum/lavka/internal/errs"
	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type ordersRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func newOrdersRepository(db *sqlx.DB, log *zap.Logger) *ordersRepository {
	return &ordersRepository{
		db:  db,
		log: log.Named("repo-orders"),
	}
}

func (r *ordersRepository) CreateOrders(ctx context.Context, orders models.OrderSlice,
) (models.OrderSlice, error) {
	cols := []string{
		models.OrderColumns.Weight, models.OrderColumns.Region,
		models.OrderColumns.DeliveryHours, models.OrderColumns.Cost}
	q := qb.Insert(schema + models.TableNames.Orders).Columns(cols...)
	for _, or := range orders {
		q = q.Values(or.Weight, or.Region, or.DeliveryHours, or.Cost)
	}
	query, args, err := q.Suffix("RETURNING *").ToSql()
	if err != nil {
		r.log.Error("build query",
			zap.Error(err),
			zap.String("query", query),
			zap.Any("args", args))
		return nil, err
	}
	ret := make(models.OrderSlice, 0, len(orders))
	if err := queries.Raw(query, args...).
		Bind(ctx, r.db, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (r *ordersRepository) GetOrder(ctx context.Context, id int) (*models.Order, error) {
	ord, err := models.FindOrder(ctx, r.db, int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return ord, nil
}

func (r *ordersRepository) ListOrders(ctx context.Context, pg model.Pagination) (models.OrderSlice, error) {
	orders, err := models.Orders(
		qm.OrderBy(models.OrderColumns.ID),
		qm.Offset(pg.Offset),
		qm.Limit(pg.Limit)).
		All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *ordersRepository) CompleteOrders(ctx context.Context, params []model.CompleteInfo) (models.OrderSlice, error) {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return nil, err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			r.log.Debug(
				"rollback error on 'CompleteOrders' method",
				zap.Error(err),
			)
		}
	}()
	orders := make(models.OrderSlice, 0, len(params))
	for _, p := range params {
		co := &models.CouriersOrder{
			OrderID:   p.OrderID,
			CourierID: p.CourierID,
		}
		exists, err := models.CouriersOrders(
			models.CouriersOrderWhere.CourierID.EQ(p.CourierID),
			models.CouriersOrderWhere.OrderID.EQ(p.OrderID),
		).Exists(ctx, tx)
		if err != nil {
			return nil, err
		}
		if !exists {
			if err := r.insertCouriersOrder(ctx, tx, co); err != nil {
				return nil, err
			}
			if _, err = models.Orders(
				models.OrderWhere.ID.EQ(p.OrderID),
				models.OrderWhere.CompleteAt.IsNull(),
			).UpdateAll(ctx, tx,
				models.M{models.OrderColumns.CompleteAt: p.CompleteTime}); err != nil {
				return nil, err
			}
		}
		ord, err := models.FindOrder(ctx, tx, p.OrderID)
		if err != nil {
			return nil, err
		}
		orders = append(orders, ord)
	}
	return orders, tx.Commit()
}

func (r *ordersRepository) insertCouriersOrder(
	ctx context.Context, exec boil.ContextExecutor, co *models.CouriersOrder) error {
	_, err := models.Orders(
		models.OrderWhere.ID.EQ(co.OrderID),
		models.OrderWhere.CompleteAt.IsNull(),
	).One(ctx, exec)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrNotFound
		}
		return err
	}
	exists, err := models.CourierExists(ctx, exec, co.CourierID)
	if err != nil {
		return err
	}
	if !exists {
		return errs.ErrNotFound
	}
	if err := co.Insert(ctx, exec, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func (r *ordersRepository) PreAssignOrders(ctx context.Context, date time.Time) ([]model.AssignCouriers, error) {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return nil, err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			r.log.Debug(
				"rollback error on 'AssignOrders' method",
				zap.Error(err),
			)
		}
	}()

	query, args, err := qb.Select("c.id as id, array_agg(o.id) as ids").
		From(schema+models.TableNames.Orders+" o, "+schema+models.TableNames.Couriers+" c").
		Join(schema+models.TableNames.AssignLim+" a on c.courier_type = a.courier_type").
		Where("o.weight <= cast(a.lim->'WeightLimit'->'Weight' as integer)").
		Where("o.region = any(c.regions)").
		Where("o."+models.OrderColumns.CompleteAt+" is null").
		Where("date(o."+models.OrderColumns.CreatedAt+") = ?", date).
		GroupBy("c." + models.CourierColumns.ID).
		ToSql()
	if err != nil {
		r.log.Error("build query",
			zap.Error(err),
			zap.String("query", query),
			zap.Any("args", args))
		return nil, err
	}

	r.log.Debug("build query",
		zap.String("query", query),
		zap.Any("args", args))
	gg := make([]groupOrders, 0)
	if err := queries.Raw(query, args...).Bind(ctx, tx, &gg); err != nil {
		return nil, err
	}

	res, err := r.fetchGroupOrders(ctx, gg, tx)
	if err != nil {
		return nil, err
	}
	return res, tx.Commit()
}

func (r *ordersRepository) fetchGroupOrders(
	ctx context.Context,
	gg []groupOrders,
	exec boil.ContextExecutor) ([]model.AssignCouriers, error) {
	orderIDs := make(map[int64]*models.Order, 0)
	for i := range gg {
		for _, id := range gg[i].OrdersIDs {
			orderIDs[id] = nil
		}
	}
	ids := make([]int64, 0, len(orderIDs))
	for id := range orderIDs {
		ids = append(ids, id)
	}
	orders, err := models.Orders(models.OrderWhere.ID.IN(ids)).All(ctx, exec)
	if err != nil {
		return nil, err
	}
	for i := range orders {
		orderIDs[orders[i].ID] = orders[i]
	}

	res := make([]model.AssignCouriers, 0)
	for i := range gg {
		courier, err := models.FindCourier(ctx, exec, gg[i].CourierID)
		if err != nil {
			return nil, err
		}
		res = append(res, model.AssignCouriers{
			Courier: courier,
			Orders: &model.AssignOrders{
				Orders: functools.MapSlice(gg[i].OrdersIDs, func(id int64) *models.Order {
					return orderIDs[id]
				}),
			},
		})
	}
	return res, nil
}

func (r *ordersRepository) StoreAssignOrders(ctx context.Context,
	ass []model.AssignCouriers, date time.Time) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			r.log.Debug(
				"rollback error on 'StoreAssignOrders' method",
				zap.Error(err),
			)
		}
	}()
	adate := &models.AssignmentsDate{
		Date: date,
	}
	if err := adate.Insert(ctx, tx, boil.Infer()); err != nil {
		return err
	}

	if len(ass) > 0 {
		cols := []string{
			models.AssignmentColumns.DateID,
			models.AssignmentColumns.CourierID,
			models.AssignmentColumns.GroupOrderIds}
		q := qb.Insert(schema + models.TableNames.Assignments).Columns(cols...)
		for _, as := range ass {
			orderIDs := functools.MapSlice(as.Orders.Orders, func(t *models.Order) int64 {
				return t.ID
			})
			q = q.Values(adate.ID, as.Courier.ID, orderIDs)
		}
		query, args, err := q.ToSql()
		if err != nil {
			r.log.Error("build query",
				zap.Error(err),
				zap.String("query", query),
				zap.Any("args", args))
			return err
		}
		r.log.Debug("build query",
			zap.String("query", query),
			zap.Any("args", args))
		if _, err = queries.Raw(query, args...).ExecContext(ctx, tx); err != nil {
			return err
		}
	}
	return tx.Commit()
}

type groupOrders struct {
	CourierID int64            `db:"id" boil:"id"`
	OrdersIDs types.Int64Array `db:"ids" boil:"ids"`
}
