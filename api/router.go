package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	*mux.Router
}

type Routing struct {
	Version    string
	PathPrefix string
	Routes     []Route
	Api        *Api
}

type Route struct {
	Version     string
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func (a *Api) newRouter() *Router {
	r := &Router{
		Router: mux.NewRouter(),
	}

	// API-Routes initialize
	apiV1 := a.newRoutingV1()
	rV1 := r.PathPrefix(apiV1.PathPrefix).Subrouter()
	for _, route := range apiV1.Routes {
		rV1.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method, "OPTIONS")
	}

	// Frontend
	r.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui/dist/"))).Methods("GET")

	return r
}
