package advisor

import "encoding/json"

func UnmarshalForecast(data []byte) (Forecast, error) {
	var r Forecast
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Forecast) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Forecast struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	State   string         `json:"state"`
	Country string         `json:"country"`
	Data    []ForecastData `json:"data"`
}

type ForecastData struct {
	Date        string      `json:"date"`
	DateBr      string      `json:"date_br"`
	Rain        Rain        `json:"rain"`
	Wind        Wind        `json:"wind"`
	Temperature Temperature `json:"temperature"`
}

type Rain struct {
	Precipitation float64 `json:"precipitation"`
}

type Temperature struct {
	Temperature int `json:"temperature"`
}

type Wind struct {
	Velocity         float64 `json:"velocity"`
	Direction        string  `json:"direction"`
	Directiondegrees float64 `json:"directiondegrees"`
	Gust             float64 `json:"gust"`
}
