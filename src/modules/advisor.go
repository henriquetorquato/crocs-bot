package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"

	types "../types"
	utils "../utils"
)

const advisorLocationBase = "http://apiadvisor.climatempo.com.br/api/v1"

// Advisor exported interface
type Advisor struct{}

// GetForecast gets a forecast for the next 'timespace' hours in the 'localeID' location
func (a Advisor) GetForecast(localeID int, timespace int, token string) types.Forecast {

	location := fmt.Sprintf("%s/forecast/locale/%d/hours/%d?token=%s", advisorLocationBase, localeID, timespace, token)
	response, err := http.Get(location)
	utils.HandleError(err)

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	forecast, err := types.UnmarshalAdvisorForecast(bodyBytes)

	return forecast
}
