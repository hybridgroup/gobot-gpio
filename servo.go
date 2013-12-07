package gobotGPIO

import (
	"github.com/hybridgroup/gobot"
	"reflect"
)

type Servo struct {
	gobot.Driver
	Adaptor      *interface{}
	CurrentAngle int
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
	gobot.Call(reflect.ValueOf(s.Adaptor).Elem().Interface(), "ServoInit")
}

func (s *Servo) Move(angle int) {
	if !(angle >= 0 && angle <= 180) {
		panic("Servo angle must be an integer between 0-180")
	}
	s.CurrentAngle = angle
	a := s.angleToSpan(angle)
	gobot.Call(reflect.ValueOf(s.Adaptor).Elem().Interface(), "ServoUpdateLocation", a, a)
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

func (s *Servo) angleToSpan(angle int) uint8 {
	return uint8(angle * (255 / 180))
}
