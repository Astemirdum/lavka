package service

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/internal/repository"
	"github.com/Astemirdum/lavka/models/v1"
	"github.com/Astemirdum/lavka/pkg/pool"
)

func (r *ordersService) assignOrders(list []model.AssignCouriers) []model.AssignCouriers {
	allocated := make(map[int64]bool)
	assigns := make([]model.AssignCouriers, 0, len(list))

	sort.Slice(list, func(i, j int) bool {
		return model.OrderCourierType(list[i].Courier.CourierType) >
			model.OrderCourierType(list[j].Courier.CourierType)
	})
	for i := range list {
		border := r.assignBorder[list[i].Courier.CourierType]
		orders := list[i].Orders.Orders
		sort.Slice(orders, func(_i, _j int) bool {
			return orders[_i].Weight < orders[_j].Weight
		})
		w := border.WeightLimit.Weight
		c := border.WeightLimit.Count
		deliver := border.DeliveryTimeRegion.First
		ords := make(models.OrderSlice, 0)
		buf := timePool.Get()
		fillDeliverTime(buf, list[i].Courier.WorkingHours)
		for _, ord := range list[i].Orders.Orders {
			if !allocated[ord.ID] &&
				(w >= ord.Weight && c >= len(ords)) &&
				timePeriodOverlap(
					buf,
					ord.DeliveryHours,
					deliver, len(ords) == 0) {
				allocated[ord.ID] = true
				ords = append(ords, ord)
				w -= ord.Weight
				c--
			}
		}
		timePool.Put(buf)
		assigns = append(assigns, model.AssignCouriers{
			Courier: list[i].Courier,
			Orders: &model.AssignOrders{
				Orders: ords,
			},
		})
	}
	return assigns
}

var timePool = pool.NewTimePool()

func timePeriodOverlap(
	buf *pool.Buf, order []string, deliver repository.DeliveryTime, first bool) bool {
	dur := deliver.First
	if !first {
		dur = deliver.Next
	}
	or := order[0]
	sp := strings.Split(or, "-")
	l, r := parseTime(sp[0]), parseTime(sp[1])
	k := (r - l) / dur
	for i := 0; i < k; i++ {
		r = l + dur
		ll := l
		segment := func() bool {
			for l <= r {
				if buf[l] != 1 {
					return false
				}
				l++
			}
			return true
		}
		if segment() {
			for ll <= r {
				buf[l] = 0
				ll++
			}
			return true
		}
		l += dur
	}
	return false
}

func fillDeliverTime(buf *pool.Buf, courHW []string) {
	for _, wh := range courHW {
		sp := strings.Split(wh, "-")
		l, r := parseTime(sp[0]), parseTime(sp[1])
		for l <= r {
			buf[l] = 1
			l++
		}
	}
}

func parseTime(s string) int {
	sp := strings.Split(s, ":")
	h, m := sp[0], sp[1]
	hh, _ := strconv.Atoi(h) //nolint:errcheck
	mm, _ := strconv.Atoi(m) //nolint:errcheck
	return hh*60 + mm
}
