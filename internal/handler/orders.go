package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Astemirdum/lavka/internal/handler/convert"

	"github.com/Astemirdum/lavka/pkg/functools"

	"github.com/Astemirdum/lavka/internal/errs"
	"github.com/Astemirdum/lavka/internal/handler/transfer"
	"github.com/Astemirdum/lavka/internal/model"
	"github.com/Astemirdum/lavka/models/v1"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateOrders(c echo.Context) error {
	var req transfer.CreateOrdersRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	orders, err := h.svc.CreateOrders(ctx, convert.ToCreateOrdersModel(&req))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &transfer.CreateOrdersResponse{Orders: convert.ToOrdersResponse(orders)})
}

func (h *Handler) GetOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	order, err := h.svc.GetOrder(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, convert.ToOrdersResponse(models.OrderSlice{order}))
}

func (h *Handler) ListOrders(c echo.Context) error {
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
	list, err := h.svc.ListOrders(c.Request().Context(), page)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &transfer.ListOrderResponse{
		Orders: convert.ToOrdersResponse(list),
	})
}

func (h *Handler) CompleteOrders(c echo.Context) error {
	var req transfer.CompleteOrdersRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	infoParams := functools.MapSlice(req.CompleteInfo,
		func(t transfer.CompleteInfo) model.CompleteInfo { return model.CompleteInfo(t) })

	orders, err := h.svc.CompleteOrders(ctx, infoParams)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(&transfer.CompleteOrdersResponse{
		Orders: convert.ToOrdersResponse(orders),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.idempotent.setCached(http.StatusOK, c.Request().Header.Clone(), body)
	return c.JSONBlob(http.StatusOK, body)
}

func (h *Handler) AssignOrders(c echo.Context) error {
	date, err := queryDateParam(c, "date")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	list, err := h.svc.AssignOrders(ctx, date)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &transfer.AssignOrdersResponse{
		Date:     date.Format(time.DateOnly),
		Couriers: convert.ToAssignResponse(list),
	})
}
