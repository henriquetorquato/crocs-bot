package utils

import (
	"time"

	types "../types"
)

const dateFormat = "2006-01-02 15:04:05"
const locationCode = "America/Sao_Paulo"

const maxPrecipitation = 0.7
const minTemperature = 22

type CrocsUse int

const (
	Use          CrocsUse = 0
	UseWithSocks CrocsUse = 1
	DontUse      CrocsUse = 2
)

func IsCrocsUsable(data []types.ForecastData) CrocsUse {

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

func FilterPeriod(timePeriod int, data []types.ForecastData) []types.ForecastData {

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
