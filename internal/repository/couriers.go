package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Astemirdum/lavka/internal/errs"
	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type couriersRepository struct {
	db                 *sqlx.DB
	log                *zap.Logger
	couriersCoefMapper couriersCoefMapper
}

type couriersCoefMapper map[models.CourierType]courierEarningsRating

func newCouriersRepository(db *sqlx.DB, log *zap.Logger) *couriersRepository {
	return &couriersRepository{
		db:                 db,
		log:                log.Named("repo-couriers"),
		couriersCoefMapper: newCouriersCoefMapper(),
	}
}

func (r *couriersRepository) CreateCouriers(ctx context.Context, couriers models.CourierSlice,
) (models.CourierSlice, error) {
	cols := []string{models.CourierColumns.CourierType, models.CourierColumns.Regions,
		models.CourierColumns.WorkingHours}
	q := qb.Insert(schema + models.TableNames.Couriers).Columns(cols...)
	for _, c := range couriers {
		q = q.Values(c.CourierType, c.Regions, c.WorkingHours)
	}
	query, args, err := q.Suffix("RETURNING *").ToSql()
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

	ret := make(models.CourierSlice, 0, len(couriers))
	if err := queries.Raw(query, args...).
		Bind(ctx, r.db, &ret); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			const checkConstraintCode = "23514"
			if pgErr.Code == checkConstraintCode {
				return nil, fmt.Errorf("%w: %s", errs.ErrCheckConstraint, pgErr.Message)
			}
		}
		return nil, err
	}
	return ret, nil
}

func (r *couriersRepository) GetCourier(ctx context.Context, id int) (*models.Courier, error) {
	cur, err := models.FindCourier(ctx, r.db, int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return cur, nil
}

func (r *couriersRepository) ListCouriers(
	ctx context.Context,
	pg model.Pagination,
) (models.CourierSlice, error) {
	list, err := models.Couriers(
		qm.OrderBy(models.CourierColumns.ID),
		qm.Offset(pg.Offset),
		qm.Limit(pg.Limit)).
		All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *couriersRepository) couriersAssignment(
	ctx context.Context,
	id int,
	date time.Time,
	exec boil.ContextExecutor,
) (models.AssignmentSlice, error) {
	exists, err := models.AssignmentsDates(
		models.AssignmentsDateWhere.Date.EQ(date),
	).Exists(ctx, exec)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errs.ErrNotFound
	}

	var lastVersion int
	if err := models.AssignmentsDates(
		qm.Select("max("+models.AssignmentsDateColumns.Version+")"),
		models.AssignmentsDateWhere.Date.EQ(date),
	).QueryRowContext(ctx, exec).Scan(&lastVersion); err != nil {
		return nil, err
	}
	qms := make([]qm.QueryMod, 0, 3)
	qms = append(qms,
		qm.InnerJoin(
			schema+models.TableNames.AssignmentsDates+" ad on ad.id = "+
				schema+models.TableNames.Assignments+"."+models.AssignmentColumns.DateID),
		qm.Where("ad."+models.AssignmentsDateColumns.Date+" = ?", date),
		qm.Where("ad."+models.AssignmentsDateColumns.Version+" = ?", lastVersion),
	)
	if id != 0 {
		qms = append(qms, models.AssignmentWhere.CourierID.EQ(int64(id)))
	}
	// ctx = boil.WithDebug(ctx, true)
	return models.Assignments(qms...).All(ctx, exec)
}

func (r *couriersRepository) GetCourierMetaInfo(
	ctx context.Context, id int64, tr model.TimeRange) (*model.CourierInfo, error) {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return nil, err
	}

	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			r.log.Debug(
				"rollback error on 'GetCourierMetaInfo' method",
				zap.Error(err),
			)
		}
	}()
	courier, err := models.FindCourier(ctx, tx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}

	earnings, err := r.calcCourierEarnings(ctx, tx, courier, tr)
	if err != nil {
		return nil, err
	}

	rating, err := r.calcCourierRating(ctx, tx, courier, tr)
	if err != nil {
		return nil, err
	}

	return &model.CourierInfo{
		Courier:  courier,
		Earnings: earnings,
		Rating:   rating,
	}, tx.Commit()
}

func (r *couriersRepository) calcCourierRating(
	ctx context.Context,
	exec boil.ContextExecutor,
	courier *models.Courier,
	tr model.TimeRange,
) (float64, error) {
	tr.To = tr.To.AddDate(0, 0, 1)
	calc := fmt.Sprintf(`%d * ( count(*) /
		(extract(EPOCH from age(timestamp '%v',
	                          	timestamp '%v'))/(60*60))
		) rating`, r.couriersCoefMapper[courier.CourierType].RatingCoef,
		tr.To.Format(time.DateTime), tr.From.Format(time.DateTime))
	query, args, err := qb.Select(calc).
		From(schema+models.TableNames.Orders+" o").
		Join(schema+models.TableNames.CouriersOrders+" co on o.id = co.order_id").
		Join(schema+models.TableNames.Couriers+" c on c.id = co.courier_id").
		Where("c."+models.CourierColumns.ID+" = ?", courier.ID).
		Where("o."+models.OrderColumns.CompleteAt+" >= ?", tr.From).
		Where("o."+models.OrderColumns.CompleteAt+" < ?", tr.To).
		ToSql()
	if err != nil {
		r.log.Error("build query",
			zap.Error(err),
			zap.String("query", query),
			zap.Any("args", args))
		return 0, err
	}

	r.log.Debug("build query",
		zap.String("query", query),
		zap.Any("args", args))
	row := queries.Raw(query, args...).QueryRowContext(ctx, exec)
	var rating float64
	if err := row.Scan(&rating); err != nil {
		return 0, err
	}
	return rating, nil
}

func (r *couriersRepository) calcCourierEarnings(
	ctx context.Context,
	exec boil.ContextExecutor,
	courier *models.Courier,
	tr model.TimeRange,
) (int, error) {
	sum := fmt.Sprintf("coalesce(sum(cost * %d), 0)",
		r.couriersCoefMapper[courier.CourierType].EarningsCoef)
	query, args, err := qb.Select(sum).
		From(schema+models.TableNames.Orders+" o").
		Join(schema+models.TableNames.CouriersOrders+" co on o.id = co.order_id").
		Join(schema+models.TableNames.Couriers+" c on c.id = co.courier_id").
		Where("c."+models.CourierColumns.ID+" = ?", courier.ID).
		Where("o."+models.OrderColumns.CompleteAt+" >= ?", tr.From).
		Where("o."+models.OrderColumns.CompleteAt+" < ?", tr.To).
		ToSql()
	if err != nil {
		r.log.Error("build query",
			zap.Error(err),
			zap.String("query", query),
			zap.Any("args", args))
		return 0, err
	}

	r.log.Debug("build query",
		zap.String("query", query),
		zap.Any("args", args))
	row := queries.Raw(query, args...).
		QueryRowContext(ctx, exec)
	var earnings int
	if err := row.Scan(&earnings); err != nil {
		return 0, err
	}
	return earnings, nil
}
