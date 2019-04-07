package main

import (
	modules "./modules"
	utils "./utils"
)

const defaultLocation = 6879
const minHours = 72
const timePeriod = 6

var platforms = []modules.Platform{
	modules.Facebook{},
	modules.Twitter{}}

func dispatchMessage(message string) {
	for _, platform := range platforms {
		platform.PostMessage(message)
	}
}

func main() {

	advisor := modules.Advisor{}
	forecast := advisor.GetForecast(defaultLocation, minHours, utils.Config.Advisor.Token)

	data := utils.FilterPeriod(timePeriod, forecast.Data)
	result := utils.IsCrocsUsable(data)
	message := utils.GetMessage(result)

	dispatchMessage(message)

}
