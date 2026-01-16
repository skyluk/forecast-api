package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/skyluk/weather-api/internal/adapters/weather"
	"github.com/skyluk/weather-api/internal/models"
)

type ApiServer struct {
	weatherAdapter weather.WeatherAdapter
}

// NewApiServer creates a new API server using the specified weather service adapter.
func NewApiServer(weatherAdapter weather.WeatherAdapter) (*ApiServer, error) {
	if weatherAdapter == nil {
		return nil, fmt.Errorf("nil weather adapter object")
	}

	return &ApiServer{weatherAdapter: weatherAdapter}, nil
}

// HandleForecastRequest handles GET requests for the forecast endpoint.
func (s *ApiServer) HandleForecastRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	coordinate := params.ByName("coordinate")

	fmt.Printf("GET /api/v1/forecast/%s\n", coordinate)

	point, err := parseCoordinate(coordinate)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	forecast, err := s.weatherAdapter.GetForecast(*point)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// write response
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Today's forecast is %s with %s temps.", forecast.Forecast, forecast.TempSummary)
}

// parseCoordinates parses latitude and longitude coordinates from a string
func parseCoordinate(coordinate string) (*models.Point, error) {
	// Assume the given lat/lon coordinates are separated by a comma
	// NOTE: in a production system I would use better validation of the input here
	elements := strings.Split(coordinate, ",")

	if len(elements) != 2 {
		return nil, fmt.Errorf("Invalid coordinate request parameter, got %s, need two comma separated vaues", coordinate)
	}

	latitude, err := strconv.ParseFloat(elements[0], 32)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for latitude from %s: %v", coordinate, err)
	}

	longitude, err := strconv.ParseFloat(elements[1], 32)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for longitude from %s: %v", coordinate, err)
	}

	return &models.Point{
		Latitude:  float32(latitude),
		Longitude: float32(longitude),
	}, nil
}

// writeErrorResponse writes details of an error response for the forecast API
func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "%v", err)
}
