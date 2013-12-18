package gobotGPIO

import (
	"github.com/hybridgroup/gobot"
	"reflect"
)

type AnalogSensor struct {
	gobot.Driver
	Adaptor *interface{}
}

func NewAnalogSensor(a interface{}) *AnalogSensor {
	b := new(AnalogSensor)
	b.Adaptor = &a
	b.Events = make(map[string]chan interface{})
	b.Commands = []string{}
	return b
}

func (a *AnalogSensor) Start() bool { return true }

func (a *AnalogSensor) Read() int {
	return int(gobot.Call(reflect.ValueOf(a.Adaptor).Elem().Interface(), "AnalogRead", a.Pin)[0].Int())
}
