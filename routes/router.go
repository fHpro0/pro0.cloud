package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	//API
	r.HandleFunc("/api/v1/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// Frontend
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(os.Getenv("frontendBuild")))).Methods("GET")

	return r
}
