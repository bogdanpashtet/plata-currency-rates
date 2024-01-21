package router

import (
	"github.com/bogdanpashtet/plata-currency-rates/internal/infrastructure/tech"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"

	_ "github.com/bogdanpashtet/plata-currency-rates/docs"
)

const (
	apiV1Prefix = "/api/v1"
)

type route struct {
	method  string
	path    string
	name    string
	handler http.HandlerFunc
}

func setRoutes(router *mux.Router, c Controller) {
	var routes = []route{
		{method: http.MethodPut, path: "", name: "UpdateRate", handler: c.UpdateRate},
		{method: http.MethodGet, path: "/by-id/{id}", name: "GetById", handler: c.GetById},
		{method: http.MethodGet, path: "/last", name: "GetLastRate", handler: c.GetLastRate},
	}

	api := router.PathPrefix(apiV1Prefix).Subrouter()

	for _, route := range routes {
		api.
			Name(route.name).
			Methods(route.method).
			Path(route.path).
			Handler(route.handler)
	}
}

func techRouter(router *mux.Router) {
	router.Methods(http.MethodGet).
		Name("prometheus").
		Path("/metrics").
		Handler(promhttp.Handler())

	router.
		PathPrefix("/tech/swagger/").
		Handler(httpSwagger.WrapHandler)

	router.Methods(http.MethodGet).
		Name("GetState").
		Path("/tech/state").
		HandlerFunc(tech.GetState)

	router.Methods(http.MethodGet).
		Name("GetInfo").
		Path("/tech/info").
		HandlerFunc(tech.GetInfo)
}
