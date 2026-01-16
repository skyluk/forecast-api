package models

import "fmt"

const (
	TempHot  = "Hot"
	TempCold = "Cold"
	TempMod  = "Moderate"
)

type SimpleForecast struct {
	Forecast    string `json:"forecast"`
	TempSummary string `json:"tempsummary"`
}

func NewSimpleForecast(forecast *WSForecastResponse) (*SimpleForecast, error) {
	if forecast == nil {
		return nil, fmt.Errorf("forecast cannot be nil")
	}

	if len(forecast.Properties.Periods) == 0 {
		return nil, fmt.Errorf("no forecast periods present")
	}

	resp := SimpleForecast{}

	for _, period := range forecast.Properties.Periods {
		if period.Num == 1 {
			resp.Forecast = period.ShortForecast
			resp.TempSummary = summarizeTemp(period.Temp)
			break
		}
	}

	return &resp, nil
}

func summarizeTemp(temp int) string {
	switch {
	case temp > 80:
		return TempHot
	case temp < 40:
		return TempCold
	default:
		return TempMod
	}
}
