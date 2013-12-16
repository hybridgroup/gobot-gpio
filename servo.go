package gobotGPIO

import (
	"github.com/hybridgroup/gobot"
	"reflect"
)

type Servo struct {
	gobot.Driver
	Adaptor      *interface{}
	CurrentAngle uint8
}

func NewServo(a interface{}) *Servo {
	s := new(Servo)
	s.CurrentAngle = 0
	s.Adaptor = &a
	s.Commands = []string{
		"MoveC",
		"MinC",
		"CenterC",
		"MaxC",
	}
	return s
}

func (s *Servo) InitServo() {
	gobot.Call(reflect.ValueOf(s.Adaptor).Elem().Interface(), "InitServo")
}

func (s *Servo) Move(angle uint8) {
	if !(angle >= 0 && angle <= 180) {
		panic("Servo angle must be an integer between 0-180")
	}
	s.CurrentAngle = angle
	gobot.Call(reflect.ValueOf(s.Adaptor).Elem().Interface(), "ServoWrite", s.Pin, s.angleToSpan(angle))
}

func (s *Servo) Min() {
	s.Move(0)
}

func (s *Servo) Center() {
	s.Move(90)
}

func (s *Servo) Max() {
	s.Move(180)
}

func (s *Servo) angleToSpan(angle uint8) uint8 {
	return uint8(angle * (255 / 180))
}
