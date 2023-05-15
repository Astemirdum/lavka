package transfer

type (
	CreateCouriersRequest struct {
		Couriers []Courier `json:"couriers" validate:"dive"`
	}
	CreateCouriersResponse struct {
		Couriers []Courier `json:"couriers"`
	}
)

type Courier struct {
	ID           int      `json:"courier_id"`
	CourierType  string   `json:"courier_type" validate:"required,oneof=FOOT BICYCLE CAR"`
	Regions      []int64  `json:"regions" validate:"required,dive,gt=0"`
	WorkingHours []string `json:"working_hours"` // HH:MM-HH:MM
}

type GetCourierResponse struct {
	Courier []Courier `json:"courier"`
}

type ListCouriersResponse struct {
	Couriers []Courier `json:"couriers"`
	Limit    int       `json:"limit"`
	Offset   int       `json:"offset"`
}

type AssignCouriersResponse struct {
	Date     string           `json:"date"`
	Couriers []AssignCouriers `json:"couriers"`
}

type MetaInfoCourierResponse struct {
	ID           int      `json:"courier_id"`
	CourierType  string   `json:"courier_type"`
	Regions      []int64  `json:"regions"`
	WorkingHours []string `json:"working_hours"`

	Rating   float64 `json:"rating,omitempty"`
	Earnings int     `json:"earnings,omitempty"`
}
