package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/skyluk/weather-api/internal/adapters/weather"
	"github.com/skyluk/weather-api/internal/server"
)

const (
	// Base URL for the weather API.
	// NOTE this would normally come from the application's configuration, not hard-coded like this.
	defaultBaseWeatherUrl = "https://api.weather.gov"
)

func main() {
	weatherAdapter, err := weather.NewWeatherAdapter(defaultBaseWeatherUrl)
	if err != nil {
		log.Fatalf("Error initializing weather adapter: %v", err)
	}

	srv, err := server.NewApiServer(weatherAdapter)
	if err != nil {
		log.Fatalf("Erorr initializing api server: %v", err)
	}

	router := httprouter.New()

	// ex. curl -X GET http://localhost:8080/api/v1/forecast/40.18443,-105.1467
	router.GET("/api/v1/forecast/:coordinate", srv.HandleForecastRequest)

	fmt.Println("Starting weather forecast API on port 8080")

	// NOTE if this was a production system, this would be initialized
	// using http.ListenAndServeTLS() or proxied via a webserver.
	http.ListenAndServe(":8080", router)
}
