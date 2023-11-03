package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"pro0.cloud/v2/routes"
)

func main() {
	err := godotenv.Load() // load .env file
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("port")
	allowOrigins := strings.Split(os.Getenv("allowOrigins"), ",")
	allowedMethods := strings.Split(os.Getenv("allowedMethods"), ",")
	allowedHeaders := strings.Split(os.Getenv("allowedHeaders"), ",")
	exposedHeaders := strings.Split(os.Getenv("exposedHeaders"), ",")

	// Start API
	log.Println("=> Starting API on port: " + port)
	apiHandlers := handlers.CORS(
		handlers.AllowedOrigins(allowOrigins),
		handlers.AllowedMethods(allowedMethods),
		handlers.AllowedHeaders(allowedHeaders),
		handlers.ExposedHeaders(exposedHeaders),
		handlers.AllowCredentials())(routes.Routes())
	_ = http.ListenAndServe(":"+port, apiHandlers)

}
