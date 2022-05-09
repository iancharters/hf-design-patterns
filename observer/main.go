package main

import (
	"github.com/iancharters/hf-design-patterns/observer/displayelement"
	"github.com/iancharters/hf-design-patterns/observer/weather"
)

func main() {
	w := weather.NewWeatherSubject()

	cc := displayelement.NewCurrentConditionsDisplay()
	w.RegisterObserver(cc)

	t := displayelement.NewTemperatureDisplay()
	w.RegisterObserver(t)

	w.SetMeasurements(weather.NewWeatherData(65.0, 92, 4))
	w.SetMeasurements(weather.NewWeatherData(41.0, 80, 3))
	w.SetMeasurements(weather.NewWeatherData(41.0, 65, 3))
	w.SetMeasurements(weather.NewWeatherData(41.0, 101, 3))
	w.SetMeasurements(weather.NewWeatherData(41.0, 45, 3))
}
