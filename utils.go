package gobotGPIO

// convert to PWM value from analog reading
func ToPwm(i int) int {
	return int((255 / 1023.0) * float64(i))
}
