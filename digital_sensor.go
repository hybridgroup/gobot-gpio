package gobotGPIO

import (
	"github.com/hybridgroup/gobot"
)

type DigitalSensorInterface interface {
	DigitalWrite(string, byte)
	DigitalRead(string) int
}

type DigitalSensor struct {
	gobot.Driver
	Adaptor DigitalSensorInterface
}

func NewDigitalSensor(a DigitalSensorInterface) *DigitalSensor {
	b := new(DigitalSensor)
	b.Adaptor = a
	b.Events = make(map[string]chan interface{})
	b.Commands = []string{
		"ReadC",
		"WriteC",
	}
	return b
}

func (a *DigitalSensor) Start() bool { return true }

func (a *DigitalSensor) Read() int {
	return a.Adaptor.DigitalRead(a.Pin)
}

func (a *DigitalSensor) Write(level byte) {
	a.Adaptor.DigitalWrite(a.Pin, level)
}
