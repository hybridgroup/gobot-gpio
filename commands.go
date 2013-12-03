package gobotGPIO

func (l *Led) ToggleC(params map[string]interface{}) {
	l.Toggle()
}
func (l *Led) OnC(params map[string]interface{}) {
	l.On()
}
func (l *Led) OffC(params map[string]interface{}) {
	l.Off()
}
