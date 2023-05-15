package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Astemirdum/lavka/models/v1"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AssignBorder map[models.CourierType]Border

func newAssignBorder(db *sqlx.DB) (AssignBorder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if _, err := db.ExecContext(ctx, "TRUNCATE TABLE "+schema+models.TableNames.AssignLim); err != nil {
		return nil, err
	}

	for cour, border := range assignBorder {
		data, err := json.Marshal(border)
		if err != nil {
			return nil, err
		}
		lim := &models.AssignLim{CourierType: cour, Lim: data}
		if err := lim.Insert(ctx, db, boil.Infer()); err != nil {
			return nil, err
		}
	}
	return assignBorder, nil
}

var assignBorder = AssignBorder{
	models.CourierTypeFOOT: {
		WeightLimit: WeightLimit{
			Weight: 10,
			Count:  2,
		},
		RegionLimit: 1,
		DeliveryTimeRegion: DeliveryTimeRegion{
			First: DeliveryTime{
				First: 25,
				Next:  10,
			},
		},
		DeliveryCost: DeliveryCost{
			First: 100,
			Next:  80,
		},
	},
	models.CourierTypeBICYCLE: {
		WeightLimit: WeightLimit{
			Weight: 20,
			Count:  4,
		},
		RegionLimit: 2,
		DeliveryTimeRegion: DeliveryTimeRegion{
			First: DeliveryTime{
				First: 12,
				Next:  8,
			},
			Next: DeliveryTime{
				First: 8,
				Next:  4,
			},
		},
		DeliveryCost: DeliveryCost{
			First: 100,
			Next:  80,
		},
	},
	models.CourierTypeCAR: {
		WeightLimit: WeightLimit{
			Weight: 40,
			Count:  7,
		},
		RegionLimit: 3,
		DeliveryTimeRegion: DeliveryTimeRegion{
			First: DeliveryTime{
				First: 8,
				Next:  4,
			},
			Next: DeliveryTime{
				First: 8,
				Next:  4,
			},
		},
		DeliveryCost: DeliveryCost{
			First: 100,
			Next:  80,
		},
	},
}

type Border struct {
	WeightLimit        WeightLimit
	RegionLimit        int
	DeliveryTimeRegion DeliveryTimeRegion
	DeliveryCost       DeliveryCost
}

type WeightLimit struct {
	Weight int
	Count  int
}

type DeliveryTime struct { // minute
	First int
	Next  int
}

type DeliveryTimeRegion struct { // minute
	First DeliveryTime
	Next  DeliveryTime
}

type DeliveryCost struct {
	First int
	Next  int // 80%
}
