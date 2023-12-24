package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Router struct {
	*mux.Router
}

type Routing struct {
	Version    string
	PathPrefix string
	Routes     []Route
	Router     *Router
}

type Route struct {
	Version     string
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func Routes() *Router {
	r := &Router{
		Router: mux.NewRouter(),
	}

	// API-Routes initialize
	apiVersions := []func() *Routing{r.newRoutingV1}
	for _, version := range apiVersions {
		v := version()
		rV := r.PathPrefix("/api" + v.PathPrefix).Subrouter()

		for _, route := range v.Routes {
			rV.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method, "OPTIONS")
		}
		log.Println("=-> initialized API route " + v.Version)
	}

	// Frontend
	r.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui/dist/"))).Methods("GET")

	return r
}
