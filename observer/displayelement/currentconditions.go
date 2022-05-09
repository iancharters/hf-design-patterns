package displayelement

import (
	"fmt"

	"github.com/iancharters/hf-design-patterns/observer/weather"
)

type CurrentConditionsDisplay struct {
	DisplayElement

	weather.WeatherData
}

func NewCurrentConditionsDisplay() *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{}
}

func (cc *CurrentConditionsDisplay) Update(wd interface{}) {
	cc.WeatherData = wd.(weather.WeatherData)
	cc.display()
}

func (cc *CurrentConditionsDisplay) display() {
	fmt.Printf("Current conditions: %.2f degrees and %.2f%% humidity\n", cc.GetTemperature(), cc.GetHumidity())
}
