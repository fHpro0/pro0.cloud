package api

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"
)

type Api struct {
	HttpServer    *http.Server
	storageSecret *rsa.PrivateKey
}

func NewApi() *Api {
	a := &Api{}

	a.newSession()

	return a
}

func (a *Api) Start(address string, handlerConfig HandlerConfig) error {
	a.HttpServer = &http.Server{
		Handler: a.newHandler(handlerConfig),
		Addr:    address,

		WriteTimeout: time.Second * 600,
		ReadTimeout:  time.Second * 600,
		IdleTimeout:  time.Second * 600,
	}

	go func() {
		// Wait for a crash at start up
		time.Sleep(1900 * time.Millisecond)
		// If not crashed, go on
		fmt.Println(fmt.Sprintf("Server is running and listening to %s using HTTPS", address))
	}()

	return a.HttpServer.ListenAndServe()
}
