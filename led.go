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
	return l
}

func (l *Led) IsOn() bool {
	return l.High
}

func (l *Led) IsOff() bool {
	return !l.IsOn()
}

func (l *Led) On() bool {
	l.changeState(l.Pin, "1")
	l.High = true
	return true
}

func (l *Led) Off() bool {
	l.changeState(l.Pin, "0")
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

func (l *Led) changeState(pin string, level string) {
	gobot.Call(reflect.ValueOf(l.Adaptor).Elem().Interface(), "DigitalWrite", pin, level)
}
