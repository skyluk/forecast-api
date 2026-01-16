package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/skyluk/forecast-api/internal/models"
)

const (
	pointPath = "/points/%f,%f"
)

// WeatherAdapter interface specification
type WeatherAdapter interface {
	GetForecast(models.Point) (*models.SimpleForecast, error)
}

type weatherAdapter struct {
	baseUrl string
}

// NewWeatherAdapter creates a new weather adapter.
func NewWeatherAdapter(baseUrl string) (WeatherAdapter, error) {
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl must not be empty")
	}

	return &weatherAdapter{
		baseUrl: baseUrl,
	}, nil
}

// GetForecast retrieves forecast details from the weather API for the given latitude/longitude and creates a SimpleForecast for and API respsonse.
func (w *weatherAdapter) GetForecast(point models.Point) (*models.SimpleForecast, error) {
	// retrieve lat/lon details for the given point to request forecast details
	p, err := w.getPointDetails(point)
	if err != nil {
		return nil, err
	}

	forecast, err := w.getForecastDetails(p)
	if err != nil {
		return nil, err
	}

	resp, err := models.NewSimpleForecast(forecast)
	if err != nil {
		return nil, fmt.Errorf("Error creating simple forecast: %v", err)
	}

	return resp, nil
}

// getPointDetails retrieves details from the weather API for the given latitude/longitude details.
func (w *weatherAdapter) getPointDetails(point models.Point) (*models.WSPointResponse, error) {
	// build the formatted path for requesting lat/lon details
	urlPath := fmt.Sprintf(pointPath, point.Latitude, point.Longitude)

	resp, err := http.Get(fmt.Sprintf("%s%s", w.baseUrl, urlPath))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var data models.WSPointResponse
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, fmt.Errorf("Error decoding JSON response: %v", err)
		}

		return &data, nil

	default:
		var data models.WSErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, fmt.Errorf("Error decoding JSON respsonse: %v", err)
		}

		return nil, fmt.Errorf("Error requesting lat/lon details: %s", data.Detail)
	}
}

// getForecastDetails gets forecast details from the weather API for the given point response.
func (w *weatherAdapter) getForecastDetails(pointResp *models.WSPointResponse) (*models.WSForecastResponse, error) {
	if pointResp.Properties.Forecast == "" {
		return nil, fmt.Errorf("Forecast URL empty")
	}

	resp, err := http.Get(pointResp.Properties.Forecast)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var data models.WSForecastResponse
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, fmt.Errorf("Error decoding JSON response: %v", err)
		}

		return &data, nil

	default:
		var data models.WSErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, fmt.Errorf("Error decoding JSON response: %v", err)
		}

		return nil, fmt.Errorf("Error requesting forecast details: %v", data.Detail)
	}
}
