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
	b.Events["update"] = make(chan interface{})
	b.Events["push"] = make(chan interface{})
	b.Events["release"] = make(chan interface{})
	b.Commands = []string{}
	return b
}

func (b *Button) StartDriver() {
	b.Active = false
	gobot.Every(b.Interval, func() {
		new_value := b.readState(b.Pin)
		if new_value != b.Active {
			b.update(new_value)
		}
	})
}

func (b *Button) readState(pin string) bool {
	i := gobot.Call(reflect.ValueOf(b.Adaptor).Elem().Interface(), "DigitalRead", pin)[0].Float()
	if i == 1 {
		return true
	} else {
		return false
	}
}

func (b *Button) update(new_val bool) {
	if new_val == true {
		b.Active = true
		b.Events["update"] <- map[string]interface{}{"push": new_val}
		b.Events["push"] <- new_val
	} else {
		b.Active = false
		b.Events["update"] <- map[string]interface{}{"release": new_val}
		b.Events["release"] <- new_val
	}
}
