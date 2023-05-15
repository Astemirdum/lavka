package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Astemirdum/lavka/models/v1"

	"github.com/Astemirdum/lavka/internal/model"

	"github.com/Astemirdum/lavka/internal/handler"
	"github.com/Astemirdum/lavka/internal/handler/convert"
	service_mocks "github.com/Astemirdum/lavka/internal/handler/mocks"
	"github.com/Astemirdum/lavka/internal/handler/transfer"
	"github.com/Astemirdum/lavka/pkg/validate"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestHandler_CreateOrders(t *testing.T) {
	t.Parallel()
	type input struct {
		body string
		req  transfer.CreateOrdersRequest
	}
	type response struct {
		expectedCode int
		expectedBody string
	}
	type mockBehavior func(r *service_mocks.MockService, req *transfer.CreateOrdersRequest)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		input        input
		response     response
		wantErr      bool
	}{
		{
			name: "ok",
			mockBehavior: func(r *service_mocks.MockService, req *transfer.CreateOrdersRequest) {
				orders := convert.ToCreateOrdersModel(req)
				for i := 0; i < len(orders); i++ {
					orders[i].ID = int64(i + 1)
				}
				r.EXPECT().
					CreateOrders(context.Background(), convert.ToCreateOrdersModel(req)).
					Return(orders, nil)
			},
			input: input{
				body: `{ "orders": [ { "weight": 5, "regions": 1, "delivery_hours": [ "12:30:15:00" ], "cost": 100 }, { "weight": 20, "regions": 2, "delivery_hours": [ "10:00:12:00" ], "cost": 230 } ] }`,
			},
			response: response{
				expectedCode: http.StatusOK,
				expectedBody: `{"orders":[{"order_id":1,"weight":5,"regions":1,"delivery_hours":["12:30:15:00"],"cost":100,"completed_time":null},{"order_id":2,"weight":20,"regions":2,"delivery_hours":["10:00:12:00"],"cost":230,"completed_time":null}]}`},
			wantErr: false,
		},
		{
			name:         "err. invalid regions 0",
			mockBehavior: func(r *service_mocks.MockService, inp *transfer.CreateOrdersRequest) {},
			input: input{
				body: `{ "orders": [ { "weight": 5, "regions": 0, "delivery_hours": [ "12:30:15:00" ], "cost": 100 } ] }`,
			},
			response: response{
				expectedCode: http.StatusBadRequest,
				expectedBody: `code=400, message=Key: 'CreateOrdersRequest.Orders[0].Region' Error:Field validation for 'Region' failed on the 'gt' tag`,
			},
			wantErr: true,
		},
		{
			name: "err. internal",
			mockBehavior: func(r *service_mocks.MockService, inp *transfer.CreateOrdersRequest) {
				r.EXPECT().CreateOrders(context.Background(), convert.ToCreateOrdersModel(inp)).
					Return(nil, errors.New("db internal"))
			},
			input: input{
				body: `{
  "orders": [
{
      "weight": 12,
      "regions": 1,
      "delivery_hours": [
        "15:00:16:00"
      ],
      "cost": 100
    }
  ]
}`,
			},
			response: response{
				expectedCode: http.StatusInternalServerError,
				expectedBody: `code=500, message=db internal`,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()
			svc := service_mocks.NewMockService(c)
			log := zap.NewExample().Named("test")
			h := handler.New(svc, log, nil)

			e := echo.New()
			e.Validator = validate.NewCustomValidator()

			r := httptest.NewRequest(
				http.MethodPost, "/orders", strings.NewReader(tt.input.body))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			ctx := e.NewContext(r, w)
			require.NoError(t, json.NewDecoder(strings.NewReader(tt.input.body)).Decode(&tt.input.req))

			tt.mockBehavior(svc, &tt.input.req)

			err := h.CreateOrders(ctx)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.response.expectedCode, w.Code)
				require.Equal(t, tt.response.expectedBody, strings.Trim(w.Body.String(), "\n"))
			} else {
				require.Error(t, err)
				er := &echo.HTTPError{}
				if errors.As(err, &er) {
					require.Equal(t, tt.response.expectedCode, er.Code)
					require.Equal(t, tt.response.expectedBody, er.Error())
				}
			}
		})
	}
}

func TestHandler_AssignOrders(t *testing.T) {
	t.Parallel()
	type input struct {
		queryParam string
		date       time.Time
	}
	type response struct {
		expectedCode int
		expectedBody string
	}
	type mockBehavior func(r *service_mocks.MockService, date time.Time)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		input        input
		response     response
		wantErr      bool
	}{
		{
			name: "ok",
			mockBehavior: func(r *service_mocks.MockService, date time.Time) {
				res := make([]model.AssignCouriers, 0)
				res = append(res, model.AssignCouriers{
					Courier: &models.Courier{ID: 1},
					Orders: &model.AssignOrders{
						Orders: models.OrderSlice{{
							ID:            1,
							Weight:        10,
							Region:        2,
							Cost:          100,
							DeliveryHours: []string{"4:30-5:30"},
						},
							{
								ID:            2,
								Weight:        5,
								Region:        2,
								Cost:          150,
								DeliveryHours: []string{"5:00-5:30"},
							},
						},
					},
				})
				r.EXPECT().
					AssignOrders(context.Background(), date).
					Return(res, nil)
			},
			input: input{
				queryParam: "date=2022-08-10",
				date:       time.Date(2022, 8, 10, 0, 0, 0, 0, time.UTC),
			},
			response: response{
				expectedCode: http.StatusOK,
				expectedBody: `{"date":"2022-08-10","couriers":[{"courier_id":1,"orders":{"group_order_id":0,"orders":[{"order_id":1,"weight":10,"regions":2,"delivery_hours":["4:30-5:30"],"cost":100,"completed_time":null},{"order_id":2,"weight":5,"regions":2,"delivery_hours":["5:00-5:30"],"cost":150,"completed_time":null}]}}]}`},
			wantErr: false,
		},
		{
			name:         "err. invalid query param",
			mockBehavior: func(r *service_mocks.MockService, date time.Time) {},
			input: input{
				queryParam: "date=2022-08-10T00:00:00",
				date:       time.Time{},
			},
			response: response{
				expectedCode: http.StatusBadRequest,
				expectedBody: `code=400, message=my own custom error for field: date values: [2022-08-10T00:00:00]`},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()
			svc := service_mocks.NewMockService(c)
			log := zap.NewExample().Named("test")
			h := handler.New(svc, log, nil)

			e := echo.New()
			e.Validator = validate.NewCustomValidator()

			r := httptest.NewRequest(http.MethodPost, "/orders/assign?"+tt.input.queryParam, nil)
			w := httptest.NewRecorder()

			ctx := e.NewContext(r, w)
			tt.mockBehavior(svc, tt.input.date)

			err := h.AssignOrders(ctx)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.response.expectedCode, w.Code)
				require.Equal(t, tt.response.expectedBody, strings.Trim(w.Body.String(), "\n"))
			} else {
				require.Error(t, err)
				er := &echo.HTTPError{}
				if errors.As(err, &er) {
					require.Equal(t, tt.response.expectedCode, er.Code)
					require.Equal(t, tt.response.expectedBody, er.Error())
				}
			}
		})
	}
}
