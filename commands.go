package gobotGPIO

// Led
func (l *Led) ToggleC(params map[string]interface{}) {
	l.Toggle()
}
func (l *Led) OnC(params map[string]interface{}) {
	l.On()
}
func (l *Led) OffC(params map[string]interface{}) {
	l.Off()
}

// Servo
func (l *Servo) MoveC(params map[string]interface{}) {
	angle := int(params["angle"].(float64))
	l.Move(angle)
}
func (l *Servo) MinC(params map[string]interface{}) {
	l.Min()
}
func (l *Servo) CenterC(params map[string]interface{}) {
	l.Center()
}
func (l *Servo) MaxC(params map[string]interface{}) {
	l.Max()
}
