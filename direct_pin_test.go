package gobotGPIO

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DirectPin", func() {
	var (
			someAdaptor TestAdaptor
			someDriver *DirectPin
		)

	BeforeEach(func() {
		someDriver = NewDirectPin(someAdaptor)
		someDriver.Pin = "1"
	})

	PIt("Should be able to DigitalRead", func() {
		Expect(true)
	})

	PIt("Should be able to DigitalWrite", func() {
		Expect(true)
	})

	PIt("Should be able to AnalogRead", func() {
		Expect(true)
	})

	PIt("Should be able to AnalogWrite", func() {
		Expect(true)
	})

	PIt("Should be able to PwmWrite", func() {
		Expect(true)
	})

	PIt("Should be able to ServoWrite", func() {
		Expect(true)
	})

	PIt("Should be able to Start", func() {
		Expect(true)
	})
})
