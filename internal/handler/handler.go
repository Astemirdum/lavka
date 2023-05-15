package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgraph-io/ristretto"

	"github.com/Astemirdum/lavka/pkg/validate"
	_ "github.com/Astemirdum/lavka/swagger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

type Handler struct {
	svc        Service
	log        *zap.Logger
	idempotent *idempotent
}

func New(srv Service, log *zap.Logger, rs *ristretto.Cache) *Handler {
	h := &Handler{
		svc:        srv,
		log:        log,
		idempotent: newIdempotent(rs, log),
	}
	return h
}

func (h *Handler) NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 4 << 10, // 4 KB
	}))
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())
	// e.Use(middleware.BodyLimit("2M"))

	const rps = 10
	e.Validator = validate.NewCustomValidator()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	api := e.Group("",
		middleware.RequestLoggerWithConfig(requestLoggerConfig()),
		middleware.RequestID())
	api.GET("/health", h.Health)
	{
		couriers := api.Group("/couriers")
		couriers.POST("", h.CreateCouriers, newRateLimiterMW(rps))
		couriers.GET("", h.ListCouriers, newRateLimiterMW(rps))
		couriers.GET("/:courier_id", h.GetCourier, newRateLimiterMW(rps))

		couriers.GET("/assignments", h.CouriersAssignments, newRateLimiterMW(rps))
		couriers.GET("/meta-info/:courier_id", h.GetCouriersMetaInfo, newRateLimiterMW(rps))
	}
	{
		orders := api.Group("/orders")
		orders.POST("", h.CreateOrders, newRateLimiterMW(rps))
		orders.GET("/:order_id", h.GetOrder, newRateLimiterMW(rps))
		orders.GET("", h.ListOrders, newRateLimiterMW(rps))
		orders.POST("/complete", h.CompleteOrders,
			newRateLimiterMW(rps),
			h.idempotent.Middleware())

		orders.POST("/assign", h.AssignOrders)
	}
	return e
}

func (h *Handler) Health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func queryDateParam(c echo.Context, query string) (time.Time, error) {
	date := time.Now().UTC()
	if err := echo.QueryParamsBinder(c).
		Time(query, &date, time.DateOnly).
		BindError(); err != nil {
		bErr := new(echo.BindingError)
		if errors.As(err, &bErr) {
			return time.Time{}, fmt.Errorf("my own custom error for field: %s values: %v", bErr.Field, bErr.Values)
		}
		return time.Time{}, err
	}
	return date, nil
}
