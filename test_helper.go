package gobotGPIO

type TestAdaptor struct {
}

func (t TestAdaptor) AnalogRead(string) int {
	return 99
}

func (t TestAdaptor) PwmWrite(string, byte) {
	return
}

func (t TestAdaptor) DigitalRead(string) int {
	return 1
}
