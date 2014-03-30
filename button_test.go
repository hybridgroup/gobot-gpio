package gobotGPIO

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Button", func() {
	var (
			someAdaptor TestAdaptor
			someDriver *Button
		)

	BeforeEach(func() {
		someDriver = NewButton(someAdaptor)
		someDriver.Pin = "1"
	})

	It("Must be able to readState", func() {
		Expect(someDriver.readState()).To(Equal(1))
	})

	PIt("Should be able to update on button state change", func() {
		Expect(true)
	})

	It("Must be able to Start", func() {
		Expect(someDriver.Start()).To(Equal(true))
	})
})
