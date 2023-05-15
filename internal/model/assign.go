package model

import "github.com/Astemirdum/lavka/models/v1"

type AssignCouriers struct {
	Courier *models.Courier
	Orders  *AssignOrders
}

type AssignOrders struct {
	GroupOrderID int64
	Orders       models.OrderSlice
}
