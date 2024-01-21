package frankfurter

import (
	"context"
	"github.com/bogdanpashtet/plata-currency-rates/internal/infrastructure/response"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFrankfurterProvider_GetRate(t *testing.T) {
	type args struct {
		ctx     context.Context
		toIso   string
		fromIso string
	}
	tests := []struct {
		name        string
		args        args
		handler     http.HandlerFunc
		want        []byte
		wantErr     bool
		wantErrText string
	}{
		{
			name: "ok test/status 200",
			args: args{
				ctx:     context.Background(),
				toIso:   "EUR",
				fromIso: "USD",
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				response.Write(w, []byte(`{"base":"EUR"}`))
			},
			want:    []byte(`{"base":"EUR"}`),
			wantErr: false,
		},
		{
			name: "fail test/status 200",
			args: args{
				ctx:     context.Background(),
				toIso:   "EUR",
				fromIso: "USD",
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "test error", http.StatusBadRequest)
			},
			wantErr:     true,
			wantErrText: "unexpected status code from provider: 400 Bad Request",
		},
		{
			name: "fail test/status 200",
			args: args{
				ctx:     context.Background(),
				toIso:   "EUR",
				fromIso: "USD",
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "test error", http.StatusInternalServerError)
			},
			wantErr:     true,
			wantErrText: "unexpected status code from provider: 500 Internal Server Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// configuration test server and test requester
			testServer := getTestServerFromHandler(tt.handler, testEndpointPath)
			defer testServer.Close()

			testRequester := getTestRequester(testServer, testEndpointPath)
			// get provider
			prv := &provider{
				getRate: testRequester,
			}
			// call method
			res, err := prv.GetRate(tt.args.ctx, tt.args.toIso, tt.args.fromIso)
			// check err and result
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.wantErrText)
			} else {
				assert.Equal(t, tt.want, res)
				assert.NoError(t, err)
			}
		})
	}
}

func TestFrankfurterProvider_GetCurrencyList(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		handler     http.HandlerFunc
		want        []byte
		wantErr     bool
		wantErrText string
	}{
		{
			name: "ok test/status 200",
			args: args{
				ctx: context.Background(),
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				response.Write(w, []byte(`{"EUR":"euro","USD":"dollar"}`))
			},
			want:    []byte(`{"EUR":"euro","USD":"dollar"}`),
			wantErr: false,
		},
		{
			name: "fail test/status 200",
			args: args{
				ctx: context.Background(),
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "test error", http.StatusBadRequest)
			},
			wantErr:     true,
			wantErrText: "unexpected status code from provider: 400 Bad Request",
		},
		{
			name: "fail test/status 200",
			args: args{
				ctx: context.Background(),
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "test error", http.StatusInternalServerError)
			},
			wantErr:     true,
			wantErrText: "unexpected status code from provider: 500 Internal Server Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// configuration test server and test requester
			testServer := getTestServerFromHandler(tt.handler, testEndpointPath)
			defer testServer.Close()

			testRequester := getTestRequester(testServer, testEndpointPath)
			// get provider
			prv := &provider{
				getCurrencyList: testRequester,
			}
			// call method
			res, err := prv.GetCurrencyList(tt.args.ctx)
			// check err and result
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.wantErrText)
			} else {
				assert.Equal(t, tt.want, res)
				assert.NoError(t, err)
			}
		})
	}
}
