package gobotGPIO

import "github.com/hybridgroup/gobot"
import "reflect"

type Led struct {
  gobot.Driver
  adaptor interface{}
  High bool
}

func NewLed(a interface{}) *Led{
  l := Led{High: false, adaptor: a}
  return &l
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
  args := []reflect.Value{reflect.ValueOf(pin), reflect.ValueOf(level)}
  reflect.ValueOf(l.adaptor).Elem().MethodByName("DigitalWrite").Call(args)
}

