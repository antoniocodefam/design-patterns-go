package main

import "fmt"

func main(){
	sportBike := &SportBike{
		name:         "Yamaha R1",
		engineCC:     998,
		topSpeed:     299,
		aerodynamics: true,
		raceMode:     true,
	}

	cruiser := &Cruiser{
		name:            "Harley-Davidson Softail",
		engineCC:        1745,
		saddlebags:      true,
		forwardControls: true,
		chromeLevel:     8,
	}

	adventureBike := &AdventureBike{
		name:            "BMW R1250GS",
		engineCC:        1254,
		offRoadTires:    true,
		luggage:         true,
		groundClearance: 220,
	}

	motorcycles := []IMotorcycle{sportBike, cruiser, adventureBike}

	fmt.Println("--- SAFETY INSPECTION ---")
	safetyInspector := &SafetyInspector{}
	for _, m := range motorcycles{
		m.inspect(safetyInspector)
	}

	fmt.Println("--- PERFORMANCE CHECK ---")
	performanceInspector := &PerformanceInspector{}
	for _, m := range motorcycles{
		m.inspect(performanceInspector)
	}
}