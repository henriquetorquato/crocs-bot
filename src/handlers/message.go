package handlers

import (
	"fmt"
	"net/http"
	"time"

	modules "../modules"
	types "../types"
	utils "../utils"
)

const dateFormat = "2006-01-02 15:04:05"
const locationCode = "America/Sao_Paulo"

const defaultLocation = 6879
const minHours = 72
const timePeriod = 6

const maxPrecipitation = 0.7
const minTemperature = 22

type CrocsUse int

const (
	Use          CrocsUse = 0
	UseWithSocks CrocsUse = 1
	DontUse      CrocsUse = 2
)

func CreatePost() {

	advisor := modules.Advisor{}
	forecast := advisor.GetForecast(defaultLocation, minHours, utils.Config.Advisor.Token)

	result := isCrocsUsable(forecast.Data, timePeriod)
	report := dispatchMessage(result)

	fmt.Println(report.ToString())

}

func dispatchMessage(crocsUse CrocsUse) types.Report {

	message := getMessage(crocsUse)
	report := types.Report{
		Time:         time.Now(),
		Message:      message,
		Publications: make([]string, 0),
		Errors:       make([]http.Header, 0)}

	for _, platform := range getPlatforms() {

		success := platform.PostMessage(message)
		if success {
			report.Publications = append(report.Publications, platform.Name())
		}

		report.Errors = utils.GetErrorHeaders()
	}

	return report
}

func getPlatforms() []types.Platform {
	return []types.Platform{
		new(modules.Facebook),
		new(modules.Twitter)}
}

func getMessage(crocsUse CrocsUse) string {
	switch crocsUse {
	case DontUse:
		return "Não."
	case UseWithSocks:
		return "Sim, com meias."
	default:
		return "Sim!"
	}
}

func isCrocsUsable(data []types.ForecastData, timePeriod int) CrocsUse {

	data = filterPeriod(timePeriod, data)

	sumTemperature := 0

	for _, forecast := range data {

		if forecast.Rain.Precipitation >= maxPrecipitation {
			return DontUse
		}

		sumTemperature += forecast.Temperature.Temperature

	}

	if (sumTemperature / len(data)) <= minTemperature {
		return UseWithSocks
	}

	return Use

}

func filterPeriod(timePeriod int, data []types.ForecastData) []types.ForecastData {

	now := time.Now()
	duration := time.Hour * time.Duration(timePeriod)
	location, _ := time.LoadLocation(locationCode)

	period := []types.ForecastData{}
	for _, forecast := range data {

		time, _ := time.ParseInLocation(dateFormat, forecast.Date, location)
		difference := time.Sub(now)

		if difference >= 0 && difference <= duration {
			period = append(period, forecast)
		}

	}

	return period

}
