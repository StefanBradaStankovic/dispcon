package dispcon

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

//
// Version v1.1.0
//
//
// To propperly define sellectedDigits[][]bool values, it is necessary to map out your display
// according to its pin numbers - eg. 1-10 for single digit screen or 1-12 for quad digit screen with sellector pins
//
// Draw an ordinal number from 0 to 9
func DisplayDrawNumber(sellectedPins []string, sellectedDigits [][]bool, number int) {

	PinsResetLow(sellectedPins[:])

	if number >= 0 && number <= 15 {
		DiodeOnClusterSellect(sellectedPins, sellectedDigits[number])
	} else {
		log.Println("Invalid number")
	}

}

// Flash every segment to visually confirm they all work
func DisplayCheckSegments(sellectedPins []string) {

	for i := 0; i < len(sellectedPins); i++ {
		DiodeFlash(sellectedPins[i], 100, 400)
	}

}

// Draw a running circle animation
func DisplayCircleRound(sellectedPins []string, speed int) {

	for i := 0; i < len(sellectedPins); i++ {
		DiodeFlash(sellectedPins[i], time.Duration(speed), 0)
	}

}

// Draw a full circle animation
func DisplayCircleFlash(sellectedPins []string, speed int) {

	for i := 0; i < len(sellectedPins); i++ {
		DiodeOn(sellectedPins[i])
	}

	time.Sleep(time.Duration(speed) * time.Millisecond)

	for i := 0; i < len(sellectedPins); i++ {
		DiodeOff(sellectedPins[i])
	}

	time.Sleep(time.Duration(speed) * time.Millisecond)

}

// Draw a full circle segment by segment
func DisplayLoading(sellectedPins []string, speed int) {

	for i := 0; i < len(sellectedPins); i++ {
		DiodeOn(sellectedPins[i])
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}

}

// Draw a startup animation
func DisplayStartupAnimation(sellectedPins []string) {

	DisplayLoading(sellectedPins, 150)
	DisplayCircleFlash(sellectedPins, 250)
	DisplayCircleFlash(sellectedPins, 250)
	DisplayCircleFlash(sellectedPins, 1500)
}

// Draw four different digits on a quad digit 8-segment display
func DisplayDrawNumberMultiple(sellectedPins []string, sellectorPins []string, sellectedDigits [][]bool, digits *[]int, mainDone *int) {

	PinsResetLow(sellectedPins[:])
	PinsResetIn(sellectorPins[:])

	for *mainDone != 1 {

		for i := 0; i < 4; i++ {
			DiodeOff(sellectorPins[i])
			DisplayDrawNumber(sellectedPins, sellectedDigits, (*digits)[i])
			time.Sleep(1 * time.Millisecond)
			DiodeIn(sellectorPins[i])
		}
	}
}

// Flash a LED by setting a pin to HIGH for timeOn amount of milliseconds with a latency of timeOff milliseconds
func DiodeFlash(inputPin string, timeOn time.Duration, timeOff time.Duration) {
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
func DiodeOnCluster(inputPins []string) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(inputPins); i++ {
		DiodeOn(inputPins[i])
	}
}

// Turn off a number of LEDs permanently by setting the corresponding pins to LOW
func DiodeOffCluster(inputPins []string) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(inputPins); i++ {
		DiodeOff(inputPins[i])
	}
}

// Turn on a number of LEDs permanently by setting the corresponding pins to HIGH
func DiodeOnClusterSellect(inputPins []string, pinState []bool) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(inputPins); i++ {
		if pinState[i] {
			DiodeOn(inputPins[i])
		}
	}
}

// Turn off a number of LEDs permanently by setting the corresponding pins to LOW
func DiodeOffClusterSellect(inputPins []string, pinState []bool) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(inputPins); i++ {
		if pinState[i] {
			DiodeOff(inputPins[i])
		}
	}
}

// Turn on an LED permanently by setting the pin to HIGH
func DiodeOn(inputPin string) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open pin for communication
	pin := gpioreg.ByName(inputPin)
	if pin == nil {
		log.Fatal("Failed to open GPIO pin")
	}

	// Set pin to IN
	if err := pin.In(gpio.PullNoChange, gpio.NoEdge); err != nil {
		log.Fatal(err)
	}

	// Set pin to HIGH
	if err := pin.Out(gpio.High); err != nil {
		log.Fatal(err)
	}
}

// Turn off an LED permanently by setting the pin to LOW
func DiodeOff(inputPin string) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open pin for communication
	pin := gpioreg.ByName(inputPin)
	if pin == nil {
		log.Fatal("Failed to open GPIO pin")
	}

	// Set pin to IN
	if err := pin.In(gpio.PullNoChange, gpio.NoEdge); err != nil {
		log.Fatal(err)
	}

	// Set pin to LOW
	if err := pin.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
}

// Completelly isolate the diode by setting the pin to IN
func DiodeIn(inputPin string) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open pin for communication
	pin := gpioreg.ByName(inputPin)
	if pin == nil {
		log.Fatal("Failed to open GPIO pin")
	}

	// Set pin to IN
	if err := pin.In(gpio.PullNoChange, gpio.NoEdge); err != nil {
		log.Fatal(err)
	}
}

// RESET ALL PINS BY SETTING THEM TO LOW
func PinsResetLow(sellectedPins []string) {

	for i := 0; i < len(sellectedPins); i++ {
		DiodeOff(sellectedPins[i])
	}

}

// RESET ALL PINS BY SETTING THEM TO IN
func PinsResetIn(sellectedPins []string) {

	for i := 0; i < len(sellectedPins); i++ {
		DiodeIn(sellectedPins[i])
	}
}
