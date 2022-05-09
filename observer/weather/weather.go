package weather

import "github.com/iancharters/hf-design-patterns/observer/observer"

type WeatherData struct {
	humidity    float32
	temperature float32
	pressure    float32
}

func NewWeatherData(humidity, temperature, pressure float32) WeatherData {
	return WeatherData{
		humidity:    humidity,
		temperature: temperature,
		pressure:    pressure,
	}
}

func (wd *WeatherData) GetTemperature() float32 {
	return wd.temperature
}

func (wd *WeatherData) GetHumidity() float32 {
	return wd.humidity
}

func (wd *WeatherData) GetPressure() float32 {
	return wd.pressure
}

type WeatherSubject struct {
	observer.Subject
	data WeatherData
}

func NewWeatherSubject() *WeatherSubject {
	return &WeatherSubject{}
}

func (ws *WeatherSubject) GetWeatherData() WeatherData {
	return ws.data
}

func (ws *WeatherSubject) SetMeasurements(wd WeatherData) {
	ws.NotifyObservers(wd)
}
