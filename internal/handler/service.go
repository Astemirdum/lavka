package handler

import (
	"context"
	"time"

	"github.com/Astemirdum/lavka/internal/service"

	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type Service interface {
	CourierService
	OrderService
}

var _ Service = (*service.Service)(nil)

type OrderService interface {
	CreateOrders(ctx context.Context, ords models.OrderSlice) (models.OrderSlice, error)
	GetOrder(ctx context.Context, id int) (*models.Order, error)
	ListOrders(ctx context.Context, pagination model.Pagination) (models.OrderSlice, error)

	CompleteOrders(ctx context.Context, params []model.CompleteInfo) (models.OrderSlice, error)
	AssignOrders(ctx context.Context, date time.Time) ([]model.AssignCouriers, error)
}

type CourierService interface {
	CreateCouriers(ctx context.Context, couriers models.CourierSlice) (models.CourierSlice, error)
	GetCourier(ctx context.Context, id int) (*models.Courier, error)
	ListCouriers(ctx context.Context, pagination model.Pagination) (models.CourierSlice, error)

	CouriersAssignment(ctx context.Context, id int, date time.Time) ([]model.AssignCouriers, error)
	GetCourierMetaInfo(ctx context.Context, id int, tr model.TimeRange) (*model.CourierInfo, error)
}
