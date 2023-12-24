package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"pro0.cloud/v2/api"
	"strings"
)

var (
	a *api.Api
)

func main() {
	err := godotenv.Load() // load .env file
	if err != nil {
		log.Fatal(err)
	}

	address := os.Getenv("address")

	handlerConfig := api.HandlerConfig{
		AllowedOrigins:   strings.Split(os.Getenv("allowOrigins"), ","),
		AllowedMethods:   strings.Split(os.Getenv("allowedMethods"), ","),
		AllowedHeaders:   strings.Split(os.Getenv("allowedHeaders"), ","),
		ExposedHeaders:   strings.Split(os.Getenv("exposedHeaders"), ","),
		AllowCredentials: true,
	}

	// Initialize server
	a = api.NewApi()

	// Start server
	_ = a.Start(address, handlerConfig)
}
