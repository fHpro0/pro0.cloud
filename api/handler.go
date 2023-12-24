package api

import (
	"github.com/gorilla/handlers"
	"net/http"
)

type Handler interface {
	http.Handler
}

type HandlerConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	AllowCredentials bool
	AccessLog        string
}

func (a *Api) newHandler(c HandlerConfig) Handler {
	allowedOrigins := handlers.AllowedOrigins(c.AllowedOrigins)
	allowedMethods := handlers.AllowedMethods(c.AllowedMethods)
	allowedHeaders := handlers.AllowedHeaders(c.AllowedHeaders)
	exposedHeaders := handlers.ExposedHeaders(c.ExposedHeaders)
	handler := handlers.CORS(
		allowedOrigins,
		allowedMethods,
		allowedHeaders,
		exposedHeaders,
		c.allowCredentials())(a.newRouter())

	return handler
}

func (c *HandlerConfig) allowCredentials() handlers.CORSOption {
	if c.AllowCredentials {
		return handlers.AllowCredentials()
	}
	return nil
}
