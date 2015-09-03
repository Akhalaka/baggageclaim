package bomberman_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"github.com/concourse/baggageclaim/bomberman"
	"github.com/concourse/baggageclaim/volume"
)

var _ = Ω

var _ = Describe("Bomberman", func() {
	It("straps a bomb to the given container with the container's grace time as the countdown", func() {
		detonated := make(chan string)

		bomberman := bomberman.New(func(handle string) {
			detonated <- handle
		})

		someVolume := volume.Volume{
			Handle: "some-handle",
			TTL:    volume.TTL(1),
		}

		bomberman.Strap(someVolume)

		select {
		case <-detonated:
		case <-time.After(someVolume.TTL.Duration() + 50*time.Millisecond):
			Fail("did not detonate!")
		}
	})

	Context("when the container has a grace time of 0", func() {
		It("never detonates", func() {

			detonated := make(chan string)

			bomberman := bomberman.New(func(handle string) {
				detonated <- handle
			})

			someVolume := volume.Volume{
				Handle: "some-handle",
				TTL:    volume.TTL(0),
			}

			bomberman.Strap(someVolume)

			select {
			case <-detonated:
				Fail("detonated!")
			case <-time.After(someVolume.TTL.Duration() + 50*time.Millisecond):
			}
		})
	})

	Describe("pausing a container's timebomb", func() {
		It("prevents it from detonating", func() {
			detonated := make(chan string)

			bomberman := bomberman.New(func(handle string) {
				detonated <- handle
			})

			someVolume := volume.Volume{
				Handle: "doomed",
				TTL:    volume.TTL(1),
			}

			bomberman.Strap(someVolume)
			bomberman.Pause("doomed")

			select {
			case <-detonated:
				Fail("detonated!")
			case <-time.After(someVolume.TTL.Duration() + 50*time.Millisecond):
			}
		})

		Context("when the handle is invalid", func() {
			It("doesn't launch any missiles or anything like that", func() {
				bomberman := bomberman.New(func(handle string) {
					panic("dont call me")
				})

				bomberman.Pause("BOOM?!")
			})
		})

		Describe("and then unpausing it", func() {
			It("causes it to detonate after the countdown", func() {

				detonated := make(chan string)

				bomberman := bomberman.New(func(handle string) {
					detonated <- handle
				})

				someVolume := volume.Volume{
					Handle: "doomed",
					TTL:    volume.TTL(1),
				}

				bomberman.Strap(someVolume)

				bomberman.Pause("doomed")

				before := time.Now()
				bomberman.Unpause("doomed")

				select {
				case <-detonated:
					Ω(time.Since(before)).Should(BeNumerically(">=", 1*time.Second))
				case <-time.After(someVolume.TTL.Duration() + 50*time.Millisecond):
					Fail("did not detonate!")
				}
			})

			Context("when the handle is invalid", func() {
				It("doesn't launch any missiles or anything like that", func() {
					bomberman := bomberman.New(func(handle string) {
						panic("dont call me")
					})

					bomberman.Unpause("BOOM?!")
				})
			})
		})
	})

	Describe("defusing a container's timebomb", func() {
		It("prevents it from detonating", func() {
			detonated := make(chan string)

			bomberman := bomberman.New(func(handle string) {
				detonated <- handle
			})

			someVolume := volume.Volume{
				Handle: "doomed",
				TTL:    volume.TTL(1),
			}

			bomberman.Strap(someVolume)
			bomberman.Defuse("doomed")

			select {
			case <-detonated:
				Fail("detonated!")
			case <-time.After(someVolume.TTL.Duration() + 50*time.Millisecond):
			}
		})

		Context("when the handle is invalid", func() {
			It("doesn't launch any missiles or anything like that", func() {
				bomberman := bomberman.New(func(handle string) {
					panic("dont call me")
				})

				bomberman.Defuse("BOOM?!")
			})
		})
	})
})
