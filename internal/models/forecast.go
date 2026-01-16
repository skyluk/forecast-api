package models

const (
	TempHot  = "Hot"
	TempCold = "Cold"
	TempMod  = "Moderate"
)

type SimpleForecast struct {
	Forecast    string `json:"forecast"`
	TempSummary string `json:"tempsummary"`
}

// summarizeTemp summarizes a temperature into the ranges "Hot", "Cold" or "Moderate"
// based on some arbitrary temperature ranges.
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
