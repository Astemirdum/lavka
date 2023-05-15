package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func TestHandler_CreateCouriers(t *testing.T) {
	t.Parallel()
	type input struct {
		body string
		req  transfer.CreateCouriersRequest
	}
	type response struct {
		expectedCode int
		expectedBody string
	}
	type mockBehavior func(r *service_mocks.MockService, req *transfer.CreateCouriersRequest)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		input        input
		response     response
		wantErr      bool
	}{
		{
			name: "ok",
			mockBehavior: func(r *service_mocks.MockService, req *transfer.CreateCouriersRequest) {
				inp := convert.ToCreateCouriersModel(req)
				for i := 0; i < len(inp); i++ {
					inp[i].ID = int64(i + 1)
				}
				r.EXPECT().
					CreateCouriers(context.Background(), convert.ToCreateCouriersModel(req)).
					Return(inp, nil)
			},
			input: input{
				body: `{
  "couriers": [
    {
      "courier_type": "FOOT",
      "regions": [
        1
      ],
      "working_hours": [
        "9:00-11:00", "14:40-17:00"
      ]
    },
{
      "courier_type": "BICYCLE",
      "regions": [
        1, 2
      ],
      "working_hours": [
        "12:00-17:00"
      ]
    }
  ]
}`,
			},
			response: response{
				expectedCode: http.StatusOK,
				expectedBody: `{"couriers":[{"courier_id":1,"courier_type":"FOOT","regions":[1],"working_hours":["9:00-11:00","14:40-17:00"]},{"courier_id":2,"courier_type":"BICYCLE","regions":[1,2],"working_hours":["12:00-17:00"]}]}`},
			wantErr: false,
		},
		{
			name:         "err. invalid regions -1",
			mockBehavior: func(r *service_mocks.MockService, inp *transfer.CreateCouriersRequest) {},
			input: input{
				body: `{"couriers":[{"courier_type":"CAR","regions":[-1, 0],"working_hours":["9:00-11:00","14:40-17:00"]}]}`,
			},
			response: response{
				expectedCode: http.StatusBadRequest,
				expectedBody: `code=400, message=Key: 'CreateCouriersRequest.Couriers[0].Regions[0]' Error:Field validation for 'Regions[0]' failed on the 'gt' tag
Key: 'CreateCouriersRequest.Couriers[0].Regions[1]' Error:Field validation for 'Regions[1]' failed on the 'gt' tag`,
			},
			wantErr: true,
		},
		{
			name: "err. internal",
			mockBehavior: func(r *service_mocks.MockService, inp *transfer.CreateCouriersRequest) {
				r.EXPECT().CreateCouriers(context.Background(), convert.ToCreateCouriersModel(inp)).
					Return(nil, errors.New("db internal"))
			},
			input: input{
				body: `{
  "couriers": [{
      "courier_type": "BICYCLE",
      "regions": [
        1, 2
      ],
      "working_hours": [
        "12:00-17:00"
      ]
    }]
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
				http.MethodPost, "/couriers", strings.NewReader(tt.input.body))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			ctx := e.NewContext(r, w)
			require.NoError(t, json.NewDecoder(strings.NewReader(tt.input.body)).Decode(&tt.input.req))

			tt.mockBehavior(svc, &tt.input.req)
			err := h.CreateCouriers(ctx)
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
