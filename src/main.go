package main

import (
	"fmt"

	advisor "./modules/advisor"
	utils "./utils"
)

const defaultLocation = 6879
const minHours = 72
const timePeriod = 6

func main() {
	forecast := advisor.GetForecast(defaultLocation, minHours, utils.Config.Advisor.Token)
	data := utils.FilterPeriod(timePeriod, forecast.Data)
	result := utils.IsCrocsUsable(data)

	fmt.Println(result)
}
