package router

import "net/http"

type Controller interface {
	UpdateRate(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	GetLastRate(w http.ResponseWriter, r *http.Request)
}
