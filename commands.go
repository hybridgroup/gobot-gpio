package gobotGPIO

import "strconv"

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
func (l *Led) BrightnessC(params map[string]interface{}) {
	level := byte(params["level"].(float64))
	l.Brightness(level)
}

// Servo
func (l *Servo) MoveC(params map[string]interface{}) {
	angle := byte(params["angle"].(float64))
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

// Digital Sensor
func (d *DigitalSensor) ReadC(params map[string]interface{}) int {
	return d.Read()
}
func (d *DigitalSensor) WriteC(params map[string]interface{}) {
	level, _ := strconv.Atoi(params["level"].(string))
	d.Write(byte(level))
}

// Analog Sensor
func (d *AnalogSensor) ReadC(params map[string]interface{}) int {
	return d.Read()
}
func (d *AnalogSensor) WriteC(params map[string]interface{}) {
	level := byte(params["level"].(float64))
	d.Write(level)
}
