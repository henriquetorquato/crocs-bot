package advisor

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const locationBase = "http://apiadvisor.climatempo.com.br/api/v1"

// GetForecast gets a forecast for the next x hours
func GetForecast(localeID int, timespace int, token string) Forecast {
	location := fmt.Sprintf("%s/forecast/locale/%d/hours/%d?token=%s", locationBase, localeID, timespace, token)
	response, err := http.Get(location)

	if (err != nil) || response.StatusCode != 200 {
		fmt.Println("KKKKK")
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	forecast, err := UnmarshalForecast(bodyBytes)

	return forecast
}
