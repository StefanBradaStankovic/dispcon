//
//  Version v2.0.1
//
//

package dispcon

import (
	"log"
	"time"
)

//  Set a single pin to INPUT
func DiodeIn(diodePinName string) {
	diodeSetState(diodePinName, 0)
}

//  Set a single pin to LOW
func DiodeOff(diodePinName string) {
	diodeSetState(diodePinName, 1)
}

//  Set a single pin to HIGH
func DiodeOn(diodePinName string) {
	diodeSetState(diodePinName, 2)
}

//  Set an array of pins to INPUT
func DiodeArrayIn(diodePinNames []string) {
	diodeArraySetState(diodePinNames, []int{}, 0)
}

//  Set an array of pins to LOW
func DiodeArrayOff(diodePinNames []string) {
	diodeArraySetState(diodePinNames, []int{}, 1)
}

//  Set an array of pins to HIGH
func DiodeArrayOn(diodePinNames []string) {
	diodeArraySetState(diodePinNames, []int{}, 2)
}

//  Flash a LED diode
//
//      - if function is called with second and third argument as 0, it will flash
//        for half a second with half a second wait time
//
func DiodeFlash(diodePinName string, timeOn time.Duration, timeOff time.Duration) {
	if timeOn == 0 {
		timeOn = 500
	}
	if timeOff == 0 {
		timeOff = 500
	}

	DiodeOn(diodePinName)
	time.Sleep(timeOn * time.Millisecond)
	DiodeOff(diodePinName)
	time.Sleep(timeOff * time.Millisecond)
}

//  Driver for displaying a symbol on a single digit screen
func DisplaySingleDigit(diodePinNames []string, diodeStates [][]int, number int) {

	DiodeArrayOff(diodePinNames)

	if number >= 0 && number <= 15 {
		diodeArraySetState(diodePinNames, diodeStates[number], 0)
	} else {
		log.Println("Invalid number")
	}

}

//  Driver for displaying symbols on a multiple digit screen
func DisplayMultipleDigits(diodePinNames []string, groundPinNames []string, diodeStates [][]int, digits *[]int, routineIsFinished *int) {

	DiodeArrayIn(diodePinNames)

	for *routineIsFinished != 1 {

		for i := 0; i < len(*digits); i++ {
			DiodeOff(groundPinNames[i])
			DisplaySingleDigit(diodePinNames, diodeStates, (*digits)[i])
			time.Sleep(1 * time.Millisecond)
			DiodeIn(groundPinNames[i])
		}
	}
}
