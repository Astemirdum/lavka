package model

type Pagination struct {
	Offset int `query:"offset" validate:"gte=0"`
	Limit  int `query:"limit" validate:"gte=0"`
}
