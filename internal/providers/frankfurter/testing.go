package frankfurter

import (
	"github.com/bogdanpashtet/plata-currency-rates/internal/infrastructure/requester"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models/config"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"time"
)

var (
	testEndpointPath = "/endpoint"
	testMethod       = http.MethodGet
	testHeaderKey    = "Test-Header"
	testHeaderValue  = "Test-Header-Value"
	testTimeout      = 2 * time.Second
)

func getTestServerFromHandler(handler http.HandlerFunc, path string) *httptest.Server {
	return newServerWithRouter(path, handler)
}

func getTestRequester(s *httptest.Server, path string) requester.Requester {
	provider := newProvider(s.URL, createTestEndpointMap(path))
	return provider
}

func createTestEndpointMap(path string) map[string]config.Endpoint {
	headers := make(http.Header)
	headers.Add(testHeaderKey, testHeaderValue)
	endpointMap := make(map[string]config.Endpoint)
	endpointMap["test-endpoint"] = config.Endpoint{
		Path:   path,
		Method: testMethod,
	}
	return endpointMap
}

func newServerWithRouter(path string, handler func(http.ResponseWriter, *http.Request)) *httptest.Server {
	r := mux.NewRouter()
	r.HandleFunc(path, handler)
	return httptest.NewServer(r)
}

func newProvider(host string, endpoints map[string]config.Endpoint) requester.Requester {
	httpClient := http.Client{Timeout: testTimeout}

	prv := config.Provider{
		Host:      host,
		Endpoints: endpoints,
	}

	return requester.New(&httpClient, prv, "test-endpoint")
}
