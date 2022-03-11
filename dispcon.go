package dispcon

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

//
// Version: v0.1.1
//

// Defining gpio.High() parameters for digits 0 1 2 3 4 5 6 7 8 9 A B C D E F
var gpioDigit = [][]bool{
	{true, true, true, true, true, true, true, false},     //	0
	{false, false, true, true, true, false, false, false}, //	1
	{true, true, false, true, true, true, false, true},    //	2
	{false, true, true, true, true, true, false, true},    //	3
	{false, false, true, true, true, false, true, true},   //	4
	{false, true, true, true, false, true, true, true},    //	5
	{true, true, true, true, false, true, true, true},     //	6
	{false, false, true, true, true, true, false, false},  //	7
	{true, true, true, true, true, true, true, true},      //	8
	{false, true, true, true, true, true, true, true},     //	9
	{true, false, true, true, true, true, true, true},     //	A
	{true, true, true, true, false, false, true, true},    //	B
	{true, true, false, true, false, true, true, false},   //	C
	{true, true, true, true, true, false, false, true},    //	D
	{true, true, false, true, false, true, true, true},    //	E
	{true, false, false, true, false, true, true, true}}   //	F

// Draw an ordinal number from 0 to 9
func DisplayDrawNumber(sellectedPins []string, number int) {

	PinsResetAll(sellectedPins[:])

	if number >= 0 && number <= 15 {
		DiodeOnCluster(sellectedPins, gpioDigit[number])
	} else {
		log.Println("Invalid number")
	}

}

//
//
// D I S P L A Y     A N I M A T I O N S
//
//

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
func DiodeOnCluster(inputPins []string, pinState []bool) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 8; i++ {
		if pinState[i] {
			DiodeOn(inputPins[i])
		}
	}
}

// Turn off a number of LEDs permanently by setting the corresponding pins to LOW
func DiodeOffCluster(inputPins []string, pinState []bool) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 8; i++ {
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

	// Set pin to LOW
	if err := pin.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}

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

	// Set pin to LOW
	if err := pin.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
}

// RESET ALL PINS BY SETTING THEM TO LOW
func PinsResetAll(sellectedPins []string) {
	for i := 0; i < len(sellectedPins); i++ {
		DiodeOff(sellectedPins[i])
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
