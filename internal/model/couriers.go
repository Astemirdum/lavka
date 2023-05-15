package model

import (
	"errors"
	"time"

	"github.com/Astemirdum/lavka/models/v1"
)

//go:generate go run github.com/dmarkham/enumer -type=CourierType -yaml -json -transform=snake -text -trimprefix=CourierType
type CourierType int

const (
	CourierTypeUnknown CourierType = iota
	CourierTypeFoot
	CourierTypeBicycle
	CourierTypeCar
)

func OrderCourierType(ct models.CourierType) CourierType {
	switch ct {
	case models.CourierTypeCAR:
		return CourierTypeCar
	case models.CourierTypeFOOT:
		return CourierTypeFoot
	case models.CourierTypeBICYCLE:
		return CourierTypeBicycle
	default:
		return CourierTypeUnknown
	}
}

type CourierInfo struct {
	Courier  *models.Courier
	Earnings int
	Rating   float64
}

type TimeRange struct {
	From time.Time `query:"startDate" validate:"required"`
	To   time.Time `query:"endDate" validate:"required"`
}

func (tr *TimeRange) Validate() error {
	if tr.From.After(tr.To) {
		return errors.New("startDate after endDate")
	}
	return nil
}
