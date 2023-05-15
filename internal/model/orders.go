package model

import (
	"time"
)

type CompleteInfo struct {
	CourierID    int64
	OrderID      int64
	CompleteTime time.Time
}
