//
//  Version v1.3.0
//
//

package dispcon

import (
	"log"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

//  Changes state of a LED diode to:
//
//	    0 - Input
//      1 - Low
//      2 - High
//
func diodeSetState(diodePinName string, diodeState int) {
	// Load all periph drivers
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Registers GPIO pin acording to its Name
	diodePin := gpioreg.ByName(diodePinName)
	if diodePin == nil {
		log.Fatal("Failed to open GPIO pin.")
	}

	if diodeState == 0 {
		if err := diodePin.In(gpio.PullNoChange, gpio.NoEdge); err != nil {
			log.Fatal(err)
		}
	}
	if diodeState == 1 {
		if err := diodePin.Out(gpio.Low); err != nil {
			log.Fatal(err)
		}
	}
	if diodeState == 2 {
		if err := diodePin.Out(gpio.High); err != nil {
			log.Fatal(err)
		}
	}
}

//  Changes state of an array of LED diodes
//
//      - to set the whole array to same state without declaring an array of ints
//        call the function with []int{} as second argument
//      - in case you are calling the function with custom diodeStates, defaultState
//        can be ignored
//
func diodeArraySetState(diodePinNames []string, diodeStates []int, defaultState int) {
	if len(diodeStates) == 0 {
		for i := 0; i < len(diodePinNames); i++ {
			diodeStates[i] = defaultState
		}
	}

	for i := 0; i < len(diodePinNames); i++ {
		diodeSetState(diodePinNames[i], diodeStates[i])
	}
}
