package dispcon

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

//
//
//

// Flash a LED by setting a pin to HIGH for timeOn amount of milliseconds with a latency of timeOff milliseconds
func diodeFlash(inputPin string, timeOn time.Duration, timeOff time.Duration) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open pin for communication
	pin := gpioreg.ByName(inputPin)
	if pin == nil {
		log.Fatal("Failed to open GPIO pin")
	}

	// Set pin to LOW
	if err := pin.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Flashing pin %s\n", inputPin)
	time.Sleep(timeOff * time.Millisecond)
	if err := pin.Out(gpio.High); err != nil {
		log.Fatal(err)
	}
	time.Sleep(timeOn * time.Millisecond)
	if err := pin.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
}

// Turn on a number of LEDs permanently by setting the corresponding pins to HIGH
func diodeOnCluster(inputPins [8]string, pinState [8]bool) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 8; i++ {
		if pinState[i] {
			diodeOn(inputPins[i])
		}
	}
}

// Turn off a number of LEDs permanently by setting the corresponding pins to LOW
func diodeOffCluster(inputPins [8]string, pinState [8]bool) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 8; i++ {
		if pinState[i] {
			diodeOff(inputPins[i])
		}
	}
}

// Turn on an LED permanently by setting the pin to HIGH
func diodeOn(inputPin string) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open pin for communication
	pin := gpioreg.ByName(inputPin)
	if pin == nil {
		log.Fatal("Failed to open GPIO pin")
	}

	// Set pin to LOW
	if err := pin.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}

	if err := pin.Out(gpio.High); err != nil {
		log.Fatal(err)
	}
}

// Turn off an LED permanently by setting the pin to LOW
func diodeOff(inputPin string) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open pin for communication
	pin := gpioreg.ByName(inputPin)
	if pin == nil {
		log.Fatal("Failed to open GPIO pin")
	}

	// Set pin to LOW
	if err := pin.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
}

// RESET ALL PINS BY SETTING THEM TO LOW
func pinsResetAll(sellectedPins []string) {
	for i := 0; i < len(sellectedPins); i++ {
		diodeOff(sellectedPins[i])
	}

	// OLD CODE 01.03.2022.
	//
	// Changed the code to be more flexible
	//
	// diodeOff(Pin01)
	// diodeOff(Pin02)
	// diodeOff(Pin04)
	// diodeOff(Pin05)
	// diodeOff(Pin06)
	// diodeOff(Pin07)
	// diodeOff(Pin09)
	// diodeOff(Pin10)
}
