package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Astemirdum/lavka/internal/handler/convert"

	"github.com/Astemirdum/lavka/internal/errs"
	"github.com/Astemirdum/lavka/internal/handler/transfer"
	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateCouriers(c echo.Context) error {
	var req transfer.CreateCouriersRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	couriers, err := h.svc.CreateCouriers(ctx, convert.ToCreateCouriersModel(&req))
	if err != nil {
		if errors.Is(err, errs.ErrCheckConstraint) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &transfer.CreateCouriersResponse{
		Couriers: convert.ToCouriersResponse(couriers),
	})
}

func (h *Handler) GetCourier(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("courier_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	courier, err := h.svc.GetCourier(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, convert.ToCouriersResponse(models.CourierSlice{courier}))
}

func (h *Handler) ListCouriers(c echo.Context) error {
	page := model.Pagination{
		Offset: 0,
		Limit:  1,
	}
	if err := c.Bind(&page); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&page); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	list, err := h.svc.ListCouriers(c.Request().Context(), page)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &transfer.ListCouriersResponse{
		Couriers: convert.ToCouriersResponse(list),
	})
}

func (h *Handler) CouriersAssignments(c echo.Context) error {
	var (
		err error
		id  int
	)
	if c.QueryParam("courier_id") != "" {
		id, err = strconv.Atoi(c.QueryParam("courier_id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}
	date, err := queryDateParam(c, "date")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	list, err := h.svc.CouriersAssignment(c.Request().Context(), id, date)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transfer.AssignCouriersResponse{
		Date:     date.Format(time.DateOnly),
		Couriers: convert.ToAssignResponse(list),
	})
}

func (h *Handler) GetCouriersMetaInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("courier_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var tr model.TimeRange
	echo.QueryParamsBinder(c).Time("startDate", &tr.From, time.DateOnly)
	echo.QueryParamsBinder(c).Time("endDate", &tr.To, time.DateOnly)
	if err := c.Validate(&tr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := tr.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	info, err := h.svc.GetCourierMetaInfo(ctx, id, tr)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &transfer.MetaInfoCourierResponse{
		ID:           int(info.Courier.ID),
		CourierType:  info.Courier.CourierType.String(),
		Regions:      info.Courier.Regions,
		WorkingHours: info.Courier.WorkingHours,
		Rating:       info.Rating,
		Earnings:     info.Earnings,
	})
}
