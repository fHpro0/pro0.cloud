package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"pro0.cloud/v2/api"
	"pro0.cloud/v2/lib/database"
	metrics "pro0.cloud/v2/lib/metric"
	"pro0.cloud/v2/lib/secureString"
	"strings"
)

var (
	a  *api.Api
	db *database.Db
	m  *metrics.Metrics
)

func main() {
	err := godotenv.Load() // load .env file
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate metrics
	m = metrics.NewMetrics()

	db, err = database.NewDb(&database.Dsn{SecureString: secureString.NewSecureString(os.Getenv("dsn"))})
	if err != nil {
		fmt.Println(fmt.Printf("database connection can`t be initialized (%s)", err))
	}
	db.Metrics = m
	fmt.Println("database initialized")

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
	a.Db = db
	a.Metrics = m
	_ = a.Start(address, handlerConfig)
}
