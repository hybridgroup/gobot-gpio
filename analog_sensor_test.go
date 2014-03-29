package gobotGPIO

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Analog-Sensor", func() {
	var (
			someAdaptor TestAdaptor
			someDriver *AnalogSensor
		)

	BeforeEach(func() {
		someDriver = NewAnalogSensor(someAdaptor)
		someDriver.Pin = "1"
	})

	It("Must be able to Read", func() {
		Expect(someDriver.Read()).To(Equal(99))
	})

	PIt("Should be able to Start", func() {
		Expect(true)
	})
})
