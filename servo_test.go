package gobotGPIO

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Servo", func() {
	var (
			someAdaptor TestAdaptor
			someDriver *Servo
		)

	BeforeEach(func() {
		someDriver = NewServo(someAdaptor)
		someDriver.Pin = "1"
	})

	PIt("Should be able to Move", func() {
		Expect(true)
	})

	PIt("Should be able to move to Min", func() {
		Expect(true)
	})

	PIt("Should be able to move to Max", func() {
		Expect(true)
	})

	PIt("Should be able to move to Center", func() {
		Expect(true)
	})

	PIt("Should be able to Start", func() {
		Expect(true)
	})
})
