package service

import (
	"context"
	"time"

	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"

	"github.com/Astemirdum/lavka/internal/repository"
	"go.uber.org/zap"
)

type ordersService struct {
	repo         OrdersRepository
	log          *zap.Logger
	assignBorder repository.AssignBorder
}

func newOrdersService(repo OrdersRepository, log *zap.Logger,
	assignBorder repository.AssignBorder) *ordersService {
	return &ordersService{
		repo:         repo,
		log:          log,
		assignBorder: assignBorder,
	}
}

func (r *ordersService) CreateOrders(ctx context.Context, orders models.OrderSlice,
) (models.OrderSlice, error) {
	if len(orders) == 0 {
		return models.OrderSlice{}, nil
	}
	return r.repo.CreateOrders(ctx, orders)
}

func (r *ordersService) GetOrder(ctx context.Context, id int) (*models.Order, error) {
	return r.repo.GetOrder(ctx, id)
}

func (r *ordersService) ListOrders(ctx context.Context, pagination model.Pagination) (models.OrderSlice, error) {
	return r.repo.ListOrders(ctx, pagination)
}

func (r *ordersService) CompleteOrders(ctx context.Context, params []model.CompleteInfo) (models.OrderSlice, error) {
	return r.repo.CompleteOrders(ctx, params)
}
func (r *ordersService) AssignOrders(ctx context.Context, date time.Time) ([]model.AssignCouriers, error) {
	list, err := r.repo.PreAssignOrders(ctx, date)
	if err != nil {
		return nil, err
	}
	assigns := r.assignOrders(list)
	if err := r.repo.StoreAssignOrders(ctx, assigns, date); err != nil {
		return nil, err
	}
	return assigns, nil
}

type OrdersRepository interface {
	CreateOrders(ctx context.Context, orders models.OrderSlice) (models.OrderSlice, error)
	GetOrder(ctx context.Context, id int) (*models.Order, error)
	ListOrders(ctx context.Context, pagination model.Pagination) (models.OrderSlice, error)
	CompleteOrders(ctx context.Context, params []model.CompleteInfo) (models.OrderSlice, error)
	PreAssignOrders(ctx context.Context, date time.Time) ([]model.AssignCouriers, error)
	StoreAssignOrders(ctx context.Context, assc []model.AssignCouriers, date time.Time) error
}

var _ OrdersRepository = (*repository.Repository)(nil)
