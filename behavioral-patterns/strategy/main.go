package main

import "fmt"

func main(){
	motorcycle := newMotorcycle("Yamaha MT-09", 60)

    printDetails(motorcycle)

	motorcycle.setRidingMode(&RainMode{})
	printDetails(motorcycle)

	motorcycle.setRidingMode(&SportMode{})
	printDetails(motorcycle)
}

func printDetails(motorcycle *motorcycle) {
	fmt.Println("Throttle response: ", motorcycle.getThrottleResponse())
	fmt.Println("ABS level: ", motorcycle.getABSLevel())
	fmt.Println("Power output: ", motorcycle.getPowerOutput())
	fmt.Println("Power output: ", motorcycle.getTractionControl())
	fmt.Println("Riding mode name: ", motorcycle.getRidingModeName())
}