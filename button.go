package gobotGPIO

import (
	"github.com/hybridgroup/gobot"
	"reflect"
)

type Button struct {
	gobot.Driver
	Adaptor *interface{}
	Active  bool
}

func NewButton(a interface{}) *Button {
	b := new(Button)
	b.Active = false
	b.Adaptor = &a
	b.Events = make(map[string]chan interface{})
	b.Events["push"] = make(chan interface{})
	b.Events["release"] = make(chan interface{})
	b.Commands = []string{}
	return b
}

func (b *Button) StartDriver() {
	state := 0
	gobot.Every(b.Interval, func() {
		new_value := b.readState(b.Pin)
		if new_value != state && new_value != -1 {
			state = new_value
			b.update(new_value)
		}
	})
}

func (b *Button) readState(pin string) int {
	return int(gobot.Call(reflect.ValueOf(b.Adaptor).Elem().Interface(), "DigitalRead", pin)[0].Int())
}

func (b *Button) update(new_val int) {
	if new_val == 1 {
		b.Active = true
		b.Events["push"] <- new_val
	} else {
		b.Active = false
		b.Events["release"] <- new_val
	}
}
