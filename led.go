package gobotGPIO

import (
	"github.com/hybridgroup/gobot"
	"reflect"
)

type Led struct {
	gobot.Driver
	Adaptor *interface{}
	High    bool
}

func NewLed(a interface{}) *Led {
	l := new(Led)
	l.High = false
	l.Adaptor = &a
	l.Commands = []string{
		"ToggleC",
		"OnC",
		"OffC",
		"BrightnessC",
	}
	return l
}

func (l *Led) IsOn() bool {
	return l.High
}

func (l *Led) IsOff() bool {
	return !l.IsOn()
}

func (l *Led) On() bool {
	l.changeState("1")
	l.High = true
	return true
}

func (l *Led) Off() bool {
	l.changeState("0")
	l.High = false
	return true
}

func (l *Led) Toggle() {
	if l.IsOn() {
		l.Off()
	} else {
		l.On()
	}
}

func (l *Led) Brightness(level uint8) {
	gobot.Call(reflect.ValueOf(l.Adaptor).Elem().Interface(), "PwmWrite", l.Pin, level)
}

func (l *Led) changeState(level string) {
	gobot.Call(reflect.ValueOf(l.Adaptor).Elem().Interface(), "DigitalWrite", l.Pin, level)
}
