package convert

import (
	"github.com/Astemirdum/lavka/internal/handler/transfer"
	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"
)

func ToAssignResponse(list []model.AssignCouriers) []transfer.AssignCouriers {
	ass := make([]transfer.AssignCouriers, 0)
	for _, cour := range list {
		ass = append(ass, transfer.AssignCouriers{
			CourierID: cour.Courier.ID,
			Orders: transfer.AssignOrders{
				GroupOrderID: int(cour.Orders.GroupOrderID),
				Orders:       ToOrdersResponse(cour.Orders.Orders),
			},
		})
	}
	return ass
}

func ToOrdersResponse(orders models.OrderSlice) []transfer.Order {
	ords := make([]transfer.Order, 0, len(orders))
	for _, order := range orders {
		ords = append(ords, transfer.Order{
			ID:            int(order.ID),
			Weight:        order.Weight,
			Region:        order.Region,
			DeliveryHours: order.DeliveryHours,
			Cost:          int(order.Cost),
			CompletedTime: order.CompleteAt,
		})
	}
	return ords
}

func ToCreateOrdersModel(req *transfer.CreateOrdersRequest) models.OrderSlice {
	couriers := make(models.OrderSlice, 0, len(req.Orders))
	for _, courier := range req.Orders {
		couriers = append(couriers, &models.Order{
			Weight:        courier.Weight,
			Region:        courier.Region,
			Cost:          int64(courier.Cost),
			DeliveryHours: courier.DeliveryHours,
		})
	}
	return couriers
}

func ToCreateCouriersModel(req *transfer.CreateCouriersRequest) []*models.Courier {
	couriers := make([]*models.Courier, 0, len(req.Couriers))
	for _, courier := range req.Couriers {
		couriers = append(couriers, &models.Courier{
			CourierType:  models.CourierType(courier.CourierType),
			Regions:      courier.Regions,
			WorkingHours: courier.WorkingHours,
		})
	}
	return couriers
}

func ToCouriersResponse(couriers models.CourierSlice) []transfer.Courier {
	list := make([]transfer.Courier, 0, len(couriers))
	for _, c := range couriers {
		list = append(list, transfer.Courier{
			ID:           int(c.ID),
			CourierType:  c.CourierType.String(),
			Regions:      c.Regions,
			WorkingHours: c.WorkingHours,
		})
	}
	return list
}
