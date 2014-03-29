package gobotGPIO

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Led", func() {
	var (
			someAdaptor TestAdaptor
			someDriver *Led
		)

	BeforeEach(func() {
		someDriver = NewLed(someAdaptor)
		someDriver.Pin = "1"
	})

	It("Must be able to tell if IsOn", func() {
		Expect(someDriver.IsOn()).To(Equal(false))
	})

	It("Must be able to tell if IsOff", func() {
		Expect(someDriver.IsOff()).To(Equal(true))
	})

	PIt("Should be able to turn On", func() {
		Expect(true)
	})

	PIt("Should be able to turn Off", func() {
		Expect(true)
	})

	PIt("Should be able to Toggle", func() {
		Expect(true)
	})

	PIt("Should be able to set Brightness", func() {
		Expect(true)
	})

	PIt("Should be able to Start", func() {
		Expect(true)
	})
})
