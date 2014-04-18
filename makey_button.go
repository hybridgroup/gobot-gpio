package gobotGPIO

import (
	"github.com/hybridgroup/gobot"
	"time"
)

type MakeyButtonInterface interface {
	DigitalRead(string) int
}

type MakeyButton struct {
	gobot.Driver
	Adaptor ButtonInterface
	Active  bool
	data    []int
}

func NewMakeyButton(a MakeyButtonInterface) *MakeyButton {
	b := new(MakeyButton)
	b.Active = false
	b.Adaptor = a
	b.Events = make(map[string]chan interface{})
	b.Events["push"] = make(chan interface{})
	b.Events["release"] = make(chan interface{})
	b.Commands = []string{}
	return b
}

func (b *MakeyButton) Start() bool {
	state := 0
	go func() {
		for {
			new_value := b.readState()
			if new_value != -1 {
				b.data = append(b.data, new_value)
				if len(b.data) > 5 {
					b.data = b.data[1:len(b.data)]
				}
				if new_value != state {
					state = new_value
					b.update(new_value)
				}
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return true
}
func (b *MakeyButton) Halt() bool { return true }
func (b *MakeyButton) Init() bool { return true }

func (b *MakeyButton) readState() int {
	return b.Adaptor.DigitalRead(b.Pin)
}

func (b *MakeyButton) update(new_val int) {
	if b.averageData() <= 0.5 {
		b.Active = true
		gobot.Publish(b.Events["push"], new_val)
	} else {
		b.Active = false
		gobot.Publish(b.Events["release"], new_val)
	}
}

func (b *MakeyButton) averageData() float64 {
	var result float64

	if len(b.data) > 0 {
		for i := range b.data {
			result += float64(b.data[i])
		}
		result = result / float64(len(b.data))
	}
	return result
}
