package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Astemirdum/lavka/pkg/functools"

	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

var qb = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type Repository struct {
	*ordersRepository
	*couriersRepository
	log          *zap.Logger
	AssignBorder AssignBorder
}

func NewRepository(db *sqlx.DB, log *zap.Logger) (*Repository, error) {
	ab, err := newAssignBorder(db)
	if err != nil {
		return nil, err
	}
	return &Repository{
		ordersRepository:   newOrdersRepository(db, log),
		couriersRepository: newCouriersRepository(db, log),
		AssignBorder:       ab,
		log:                log.Named("repo"),
	}, nil
}

func (r *Repository) CouriersAssignment(ctx context.Context, id int, date time.Time,
) ([]model.AssignCouriers, error) {
	var ac []model.AssignCouriers
	err := runTx(ctx, r.couriersRepository.db, func(tx *sqlx.Tx) error {
		list, err := r.couriersRepository.couriersAssignment(ctx, id, date, tx)
		if err != nil {
			return err
		}
		gg := functools.MapSlice(list, func(t *models.Assignment) groupOrders {
			return groupOrders{
				CourierID: t.CourierID,
				OrdersIDs: t.GroupOrderIds,
			}
		})
		ac, err = r.fetchGroupOrders(ctx, gg, tx)
		if err != nil {
			return err
		}
		return nil
	})
	return ac, err
}

const (
	schema = "lk."
)

type TxRunner interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

type TxFunc func(tx *sqlx.Tx) error

func runTx(ctx context.Context, db TxRunner, fn TxFunc) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	if err = fn(tx); err != nil {
		// rbErr := tx.Rollback(); rbErr != nil {
		//
		//}
		//!errors.Is(err, sql.ErrTxDone)
		return multierr.Combine(err, tx.Rollback())
	}
	return tx.Commit()
}

type courierEarningsRating struct {
	EarningsCoef int
	RatingCoef   int
}

func newCouriersCoefMapper() couriersCoefMapper {
	return map[models.CourierType]courierEarningsRating{
		models.CourierTypeFOOT: {
			EarningsCoef: 2,
			RatingCoef:   3,
		},
		models.CourierTypeBICYCLE: {
			EarningsCoef: 3,
			RatingCoef:   2,
		},
		models.CourierTypeCAR: {
			EarningsCoef: 4,
			RatingCoef:   1,
		},
	}
}
