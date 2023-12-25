package api

import "net/http"

type RoutingV1 struct {
	Routing
}

func (a *Api) newRoutingV1() *RoutingV1 {
	routing := &RoutingV1{
		Routing{Api: a},
	}
	routing.Version = "v1"
	routing.PathPrefix = "/" + routing.Version
	routing.Routes = routing.getV1Routes()

	return routing
}

func (r *RoutingV1) getV1Routes() []Route {

	return []Route{
		{
			Version: r.Version,
			Name:    "Status",
			Method:  "GET",
			Pattern: "/status",
			Public:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("ok"))
			},
		},
		{
			Version: r.Version,
			Name:    "Status",
			Method:  "GET",
			Pattern: "/admin",
			Public:  false,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("ok"))
			},
		},
	}
}
