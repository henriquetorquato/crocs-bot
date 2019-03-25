package main

import (
	advisor "./modules/advisor"
	facebook "./modules/facebook"
	utils "./utils"
)

const defaultLocation = 6879
const minHours = 72
const timePeriod = 6

func main() {
	forecast := advisor.GetForecast(defaultLocation, minHours, utils.Config.Advisor.Token)
	data := utils.FilterPeriod(timePeriod, forecast.Data)
	result := utils.IsCrocsUsable(data)
	message := utils.GetMessage(result)
	facebook.PostMessage(message, utils.Config.Facebook.Page.ID, utils.Config.Facebook.Page.Token)
}
