package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	modules "../modules"
	facebookModule "../modules/facebook"
	twitterModule "../modules/twitter"
	types "../types"
	utils "../utils"
)

const dateFormat = "2006-01-02 15:04:05"
const locationCode = "America/Sao_Paulo"

// Belo Horizonte
const defaultLocation = 6879

// Advisor API only returns forecast for the next 72h minimum
const minHours = 72

const maxPrecipitation = 0.7
const minTemperature = 22

type CrocsUse int

const (
	Use          CrocsUse = 0
	UseWithSocks CrocsUse = 1
	DontUse      CrocsUse = 2
)

func CreatePost(period int) {

	advisor := modules.Advisor{}
	forecast := advisor.GetForecast(defaultLocation, minHours, utils.Config.Advisor.Token)

	message := buildMessage(
		isCrocsUsable(forecast.Data, period))

	report := dispatchMessage(message)

	fmt.Println(report.ToString())

}

func dispatchMessage(message string) types.Report {

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
		new(facebookModule.Facebook),
		new(twitterModule.Twitter)}
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

func buildMessage(crocsUse CrocsUse) string {

	var message string
	var variation []string

	messages := utils.Config.Message
	variations := utils.Config.Variations

	switch crocsUse {
	case DontUse:
		message = messages.DontUse
		variation = variations.Rain
	case UseWithSocks:
		message = messages.UseWithSocks
		variation = variations.Cold
	default:
		message = messages.Use
		variation = variations.Heat
	}

	lastPosts := make([]string, 0)
	for _, platform := range getPlatforms() {
		lastPost := platform.GetLastPost()
		lastPosts = append(lastPosts, lastPost.Content)
	}

	if utils.Contains(lastPosts, message) {

		message = fmt.Sprintf("%s, %s", message,
			variation[rand.Intn(len(variation))])

	}

	return message

}
