package displayelement

import (
	"fmt"

	"github.com/iancharters/hf-design-patterns/observer/weather"
)

type TemperatureDisplay struct {
	DisplayElement

	historicalTemperatures []float32
}

func NewTemperatureDisplay() *TemperatureDisplay {
	return &TemperatureDisplay{}
}

func (cc *TemperatureDisplay) Update(wd interface{}) {
	if wd, ok := wd.(weather.WeatherData); ok {
		cc.historicalTemperatures = append(cc.historicalTemperatures, wd.GetTemperature())
		cc.display()
	}
}

func (cc *TemperatureDisplay) display() {
	fmt.Printf(
		"Avg/Max/Min temperature: %.1f/%.1f/%.1f\n",
		cc.getAverageTemperature(),
		cc.getMaxTemperature(),
		cc.getMinTemperature(),
	)
}

func (cc *TemperatureDisplay) getAverageTemperature() float32 {
	var (
		sum float32
		n   = len(cc.historicalTemperatures)
	)

	for _, m := range cc.historicalTemperatures {
		sum += m
	}

	return sum / float32(n)

}

func (cc *TemperatureDisplay) getMaxTemperature() float32 {
	var max float32 = cc.historicalTemperatures[0]

	for _, t := range cc.historicalTemperatures {
		if t > max {
			max = t
		}
	}

	return max
}

func (cc *TemperatureDisplay) getMinTemperature() float32 {
	var min float32 = cc.historicalTemperatures[0]

	for _, t := range cc.historicalTemperatures {
		if t < min {
			min = t
		}
	}

	return min
}
