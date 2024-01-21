package controller

import (
	"bytes"
	"errors"
	"github.com/bogdanpashtet/plata-currency-rates/internal/controller/mocks"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_controller_GetByWbTrackUuid(t *testing.T) {
	testReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1?rate=EUR/USD", strings.NewReader(""))

	testReq2 := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1", strings.NewReader(""))

	testReq3 := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/?rate=EUR/RUB", strings.NewReader(""))

	updResp := models.UpdateResponse{RateId: "dsfsdfsdfsdfsdf"}

	type fields struct {
		service       *mocks.MockService
		validIsoCodes map[string]struct{}
		logger        zerolog.Logger
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    *httptest.ResponseRecorder
	}{
		{
			name: "ok test. valid input, validation ok",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetRateFromProvider(gomock.Any(), "EUR", "USD").
						Return(updResp, nil),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      200,
				Body:      bytes.NewBuffer([]byte((`{"rateId":"dsfsdfsdfsdfsdf"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "ok test.  validation false empty query parameter",
			prepare: func(f *fields) {
				gomock.InOrder()
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq2,
			},
			want: &httptest.ResponseRecorder{
				Code:      400,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Bad Request","status":400,"detail":"parameter doesn't match pattern EUR/USD"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test. validation fail iso code doesn't exist",
			prepare: func(f *fields) {
				gomock.InOrder()
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq3,
			},
			want: &httptest.ResponseRecorder{
				Code:      400,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Bad Request","status":400,"detail":"uexpected iso code RUB. try this one: EUR USD"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test.  validation fail iso code doesn't exist",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetRateFromProvider(gomock.Any(), "EUR", "USD").
						Return(updResp, errors.New("test-error")),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      500,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Internal Server Error","status":500,"detail":"test-error"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test.  validation fail iso code doesn't exist",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetRateFromProvider(gomock.Any(), "EUR", "USD").
						Return(updResp, errors.New("test-error")),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      500,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Internal Server Error","status":500,"detail":"test-error"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			defer mc.Finish()
			f := fields{
				service:       mocks.NewMockService(mc),
				validIsoCodes: map[string]struct{}{"EUR": {}, "USD": {}},
				logger:        zerolog.Nop(),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			c := New(f.service, f.validIsoCodes, f.logger)

			respWriter := httptest.NewRecorder()

			c.UpdateRate(respWriter, tt.args.req)

			assert.Equal(t, tt.want.Code, respWriter.Code)
			assert.Equal(t, tt.want.HeaderMap, respWriter.HeaderMap)
			assert.Equal(t, tt.want.Body.String(), respWriter.Body.String())
		})
	}
}

func Test_controller_GetById(t *testing.T) {
	testReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/by-id", strings.NewReader(""))
	testReq = mux.SetURLVars(testReq, map[string]string{"id": "700e7e70-ecd8-4ed3-bf60-fed516bccee8"})

	testReq2 := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/by-id", strings.NewReader(""))
	testReq2 = mux.SetURLVars(testReq2, map[string]string{"id": "fdg"})

	testReq3 := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/by-id", strings.NewReader(""))

	timeP, _ := time.Parse("2006-01-02", "2003-02-03")

	updResp := models.CurrencyRateWithDt{
		Id:       "700e7e70-ecd8-4ed3-bf60-fed516bccee8",
		Currency: "EUR",
		Base:     "USD",
		Rate:     0.11,
		UpdateDt: timeP,
	}

	type fields struct {
		service       *mocks.MockService
		validIsoCodes map[string]struct{}
		logger        zerolog.Logger
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    *httptest.ResponseRecorder
	}{
		{
			name: "ok test. valid input, validation ok",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetById(gomock.Any(), "700e7e70-ecd8-4ed3-bf60-fed516bccee8").
						Return(updResp, nil),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      200,
				Body:      bytes.NewBuffer([]byte((`{"id":"700e7e70-ecd8-4ed3-bf60-fed516bccee8","currency":"EUR","base":"USD","rate":0.11,"updateDt":"2003-02-03T00:00:00Z"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test. invalid input uuid incorrect",
			prepare: func(f *fields) {
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq2,
			},
			want: &httptest.ResponseRecorder{
				Code:      400,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Bad Request","status":400,"detail":"invalid UUID length: 3"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test. invalid input uuid isn't set",
			prepare: func(f *fields) {
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq3,
			},
			want: &httptest.ResponseRecorder{
				Code:      400,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Bad Request","status":400,"detail":"set id value"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test. valid input, service returns err",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetById(gomock.Any(), "700e7e70-ecd8-4ed3-bf60-fed516bccee8").
						Return(updResp, errors.New("test-error")),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      500,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Internal Server Error","status":500,"detail":"test-error"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test. valid input, service returns err",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetById(gomock.Any(), "700e7e70-ecd8-4ed3-bf60-fed516bccee8").
						Return(updResp, errors.New("no rows in result set")),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      500,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Internal Server Error","status":500,"detail":"no rows in result set"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			defer mc.Finish()
			f := fields{
				service:       mocks.NewMockService(mc),
				validIsoCodes: map[string]struct{}{"EUR": {}, "USD": {}},
				logger:        zerolog.Nop(),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			c := New(f.service, f.validIsoCodes, f.logger)

			respWriter := httptest.NewRecorder()

			c.GetById(respWriter, tt.args.req)

			assert.Equal(t, tt.want.Code, respWriter.Code)
			assert.Equal(t, tt.want.HeaderMap, respWriter.HeaderMap)
			assert.Equal(t, tt.want.Body.String(), respWriter.Body.String())
		})
	}
}

func Test_controller_GetLastRate(t *testing.T) {
	testReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/last?rate=EUR/USD", strings.NewReader(""))

	testReq2 := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/last", strings.NewReader(""))

	testReq3 := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/last?rate=EUR/RUB", strings.NewReader(""))

	timeP, _ := time.Parse("2006-01-02", "2003-02-03")

	updResp := models.CurrencyRateLast{
		Currency: "EUR",
		Base:     "USD",
		Rate:     0.11,
		UpdateDt: timeP,
	}

	type fields struct {
		service       *mocks.MockService
		validIsoCodes map[string]struct{}
		logger        zerolog.Logger
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    *httptest.ResponseRecorder
	}{
		{
			name: "ok test. valid input, validation ok",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetLastRate(gomock.Any(), "EUR", "USD").
						Return(updResp, nil),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      200,
				Body:      bytes.NewBuffer([]byte((`{"currency":"EUR","base":"USD","rate":0.11,"updateDt":"2003-02-03T00:00:00Z"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "ok test.  validation false empty query parameter",
			prepare: func(f *fields) {
				gomock.InOrder()
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq2,
			},
			want: &httptest.ResponseRecorder{
				Code:      400,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Bad Request","status":400,"detail":"parameter doesn't match pattern EUR/USD"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test. validation fail iso code doesn't exist",
			prepare: func(f *fields) {
				gomock.InOrder()
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq3,
			},
			want: &httptest.ResponseRecorder{
				Code:      400,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Bad Request","status":400,"detail":"uexpected iso code RUB. try this one: EUR USD"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test.  validation fail iso code doesn't exist",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetLastRate(gomock.Any(), "EUR", "USD").
						Return(updResp, errors.New("test-error")),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      500,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Internal Server Error","status":500,"detail":"test-error"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
		{
			name: "fail test.  validation fail iso code doesn't exist",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.service.EXPECT().
						GetLastRate(gomock.Any(), "EUR", "USD").
						Return(updResp, errors.New("test-error")),
				)
			},
			args: args{
				w:   httptest.NewRecorder(),
				req: testReq,
			},
			want: &httptest.ResponseRecorder{
				Code:      500,
				Body:      bytes.NewBuffer([]byte((`{"type":"about:blank","title":"Internal Server Error","status":500,"detail":"test-error"}`))),
				HeaderMap: map[string][]string{"Content-Type": {"application/json"}},
				Flushed:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			defer mc.Finish()
			f := fields{
				service:       mocks.NewMockService(mc),
				validIsoCodes: map[string]struct{}{"EUR": {}, "USD": {}},
				logger:        zerolog.Nop(),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			c := New(f.service, f.validIsoCodes, f.logger)

			respWriter := httptest.NewRecorder()

			c.GetLastRate(respWriter, tt.args.req)

			assert.Equal(t, tt.want.Code, respWriter.Code)
			assert.Equal(t, tt.want.HeaderMap, respWriter.HeaderMap)
			assert.Equal(t, tt.want.Body.String(), respWriter.Body.String())
		})
	}
}
