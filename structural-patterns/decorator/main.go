package main

import "fmt"

func main(){
	basicMotorcycle := &BaseMotorcycle{
		description: "Standard kit bike",
		price: 6200.00,
	}
    printDetails(basicMotorcycle)

	sportsBike := newAbsMotorcycleDecorator(basicMotorcycle)
	printDetails(sportsBike)

	comfortBike := newHeatedGripsMotorcycleDecorator(sportsBike)
	printDetails(comfortBike)

	travelBike := newLuggageMotorcycleDecorator(comfortBike)
	printDetails(travelBike)
}

func printDetails(m IMotorcycle) {
	fmt.Printf("Details: %s - Price: %.2f\n", m.getDescription(), m.getPrice())
}