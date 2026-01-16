package models

import "fmt"

// Point represents a latitude/longitude coordinate
type Point struct {
	Latitude  float32
	Longitude float32
}

// WSPointResponse represents basic details of the the response from the weather service's 'point' endpoint
type WSPointResponse struct {
	Id         string          `json:"id"`
	Properties WSPointProperty `json:"properties"`
}

// WSPointPoperty represents basic details of the 'properties' detail in the WSPointResponse object
type WSPointProperty struct {
	Forecast string `json:"forecast"`
}

type WSForecastResponse struct {
	Id         string
	Properties WSForecastProperty `json:"properties"`
}

type WSForecastProperty struct {
	Periods []WSForecastPeriod `json:"periods"`
}

type WSForecastPeriod struct {
	Num           int    `json:"number"`
	Temp          int    `json:"temperature"`
	ShortForecast string `json:"shortForecast"`
}

// WSErrorResponse represents basic details for an error response from the weather service
type WSErrorResponse struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func (fc *WSForecastResponse) CreateSimpleForecast() (*SimpleForecast, error) {
	if fc == nil {
		return nil, fmt.Errorf("forecast cannot be nil")
	}

	if len(fc.Properties.Periods) == 0 {
		return nil, fmt.Errorf("no forecast periods present")
	}

	resp := SimpleForecast{}

	for _, period := range fc.Properties.Periods {
		if period.Num == 1 {
			resp.Forecast = period.ShortForecast
			resp.TempSummary = summarizeTemp(period.Temp)
			break
		}
	}

	return &resp, nil
}
