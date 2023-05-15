package transfer

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type Order struct {
	ID            int       `json:"order_id"`
	Weight        int       `json:"weight"`
	Region        int       `json:"regions" validate:"gt=0"`
	DeliveryHours []string  `json:"delivery_hours"` // HH:MM-HH:MM
	Cost          int       `json:"cost"`
	CompletedTime null.Time `json:"completed_time,omitempty"`
}

type CreateOrdersRequest struct {
	Orders []Order `json:"orders" validate:"required,dive"`
}

type CreateOrdersResponse struct {
	Orders []Order `json:"orders"`
}

type ListOrderResponse struct {
	Orders []Order `json:"orders"`
}

type (
	CompleteOrdersRequest struct {
		CompleteInfo []CompleteInfo `json:"complete_info" validate:"required,dive"`
	}
	CompleteOrdersResponse struct {
		Orders []Order `json:"orders"`
	}
)

type CompleteInfo struct {
	CourierID    int64     `json:"courier_id" validate:"required"`
	OrderID      int64     `json:"order_id" validate:"required"`
	CompleteTime time.Time `json:"complete_time" validate:"required"`
}

type AssignOrdersResponse struct {
	Date     string           `json:"date"`
	Couriers []AssignCouriers `json:"couriers"`
}

type AssignCouriers struct {
	CourierID int64        `json:"courier_id"`
	Orders    AssignOrders `json:"orders"`
}

type AssignOrders struct {
	GroupOrderID int     `json:"group_order_id"`
	Orders       []Order `json:"orders"`
}
