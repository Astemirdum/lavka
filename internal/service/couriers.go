package service

import (
	"context"
	"time"

	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"

	"github.com/Astemirdum/lavka/internal/repository"
	"go.uber.org/zap"
)

type couriersService struct {
	repo CouriersRepository
	log  *zap.Logger
}

type CouriersRepository interface {
	CreateCouriers(
		ctx context.Context, couriers models.CourierSlice) (models.CourierSlice, error)
	GetCourier(ctx context.Context, id int) (*models.Courier, error)
	ListCouriers(ctx context.Context, pagination model.Pagination) (models.CourierSlice, error)

	CouriersAssignment(ctx context.Context, id int, date time.Time) ([]model.AssignCouriers, error)
	GetCourierMetaInfo(ctx context.Context, id int64, tr model.TimeRange) (*model.CourierInfo, error)
}

var _ CouriersRepository = (*repository.Repository)(nil)

func newCouriersService(repo CouriersRepository, log *zap.Logger) *couriersService {
	return &couriersService{
		repo: repo,
		log:  log,
	}
}

func (r *couriersService) CreateCouriers(ctx context.Context, couriers models.CourierSlice) (models.CourierSlice, error) {
	return r.repo.CreateCouriers(ctx, couriers)
}

func (r *couriersService) GetCourier(ctx context.Context, id int) (*models.Courier, error) {
	return r.repo.GetCourier(ctx, id)
}

func (r *couriersService) ListCouriers(ctx context.Context, pagination model.Pagination) (models.CourierSlice, error) {
	return r.repo.ListCouriers(ctx, pagination)
}

func (r *couriersService) CouriersAssignment(ctx context.Context, id int, date time.Time) ([]model.AssignCouriers, error) {
	return r.repo.CouriersAssignment(ctx, id, date)
}

func (r *couriersService) GetCourierMetaInfo(
	ctx context.Context, id int, tr model.TimeRange,
) (*model.CourierInfo, error) {
	return r.repo.GetCourierMetaInfo(ctx, int64(id), tr)
}
