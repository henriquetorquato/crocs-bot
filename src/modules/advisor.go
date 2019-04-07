package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	types "../types"
	handler "../utils/error"
)

const advisorLocationBase = "http://apiadvisor.climatempo.com.br/api/v1"

type Advisor struct{}

// GetForecast gets a forecast for the next x hours
func (a Advisor) GetForecast(localeID int, timespace int, token string) types.Forecast {

	location := fmt.Sprintf("%s/forecast/locale/%d/hours/%d?token=%s", advisorLocationBase, localeID, timespace, token)
	response, err := http.Get(location)
	handler.HandleError(err)

	// if (err != nil) || response.StatusCode != 200 {
	// 	fmt.Println("KKKKK")
	// }

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	forecast, err := unmarshalForecast(bodyBytes)

	return forecast
}

func unmarshalForecast(data []byte) (types.Forecast, error) {
	var r types.Forecast
	err := json.Unmarshal(data, &r)
	return r, err
}
