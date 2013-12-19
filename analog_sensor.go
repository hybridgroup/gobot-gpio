package gobotGPIO

import (
	"github.com/hybridgroup/gobot"
)

type AnalogSensorInterface interface {
	AnalogRead(string) int
	PwmWrite(string, uint8)
}

type AnalogSensor struct {
	gobot.Driver
	Adaptor AnalogSensorInterface
}

func NewAnalogSensor(a AnalogSensorInterface) *AnalogSensor {
	b := new(AnalogSensor)
	b.Adaptor = a
	b.Events = make(map[string]chan interface{})
	b.Commands = []string{
		"ReadC",
		"WriteC",
	}
	return b
}

func (a *AnalogSensor) Start() bool { return true }

func (a *AnalogSensor) Read() int {
	return a.Adaptor.AnalogRead(a.Pin)
}

func (a *AnalogSensor) Write(level uint8) {
	a.Adaptor.PwmWrite(a.Pin, level)
}
