package routes

import "net/http"

func (r *Router) newRoutingV1() *Routing {
	routing := &Routing{
		Version: "v1",
		Router:  r,
	}
	routing.PathPrefix = "/" + routing.Version
	routing.Routes = routing.getV1Routes()

	return routing
}

func (r *Routing) getV1Routes() []Route {

	return []Route{
		{
			Version: r.Version,
			Name:    "Status",
			Method:  "GET",
			Pattern: "/status",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("ok"))
			},
		},
	}
}
