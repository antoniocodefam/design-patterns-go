package main

import "fmt"

type PerformanceInspector struct{}

func (pi *PerformanceInspector) inspectSportBike(s *SportBike) {
	fmt.Printf("Performance check: %s\n", s.name)
	fmt.Printf("  - Engine: %dcc - Excellent for racing\n", s.engineCC)
	fmt.Printf("  - Top speed: %d km/h\n", s.topSpeed)
	if s.raceMode {
		fmt.Println("  - Race mode enabled - optimal performance")
	}
	powerToWeight := float64(s.engineCC) / 200.0
	fmt.Printf("  - Power-to-weight ratio: %.2f (higher is better)\n", powerToWeight)
	fmt.Println("  ✓ Excellent performance for track use")
}

func (pi *PerformanceInspector) inspectCruiser(c *Cruiser) {
	fmt.Printf("Performance check: %s\n", c.name)
	fmt.Printf("  - Engine: %dcc - Good torque for cruising\n", c.engineCC)
	fmt.Println("  - Optimized for comfort, not speed")
	if c.forwardControls {
		fmt.Println("  - Forward controls enhance relaxed riding")
	}
	fmt.Println("  ✓ Performance suited for highway cruising")
}

func (pi *PerformanceInspector) inspectAdventureBike(a *AdventureBike) {
	fmt.Printf("Performance check: %s\n", a.name)
	fmt.Printf("  - Engine: %dcc - Versatile power delivery\n", a.engineCC)
	fmt.Printf("  - Ground clearance: %dmm - ", a.groundClearance)
	if a.groundClearance > 200 {
		fmt.Println("Excellent for off-road")
	} else {
		fmt.Println("Adequate for light trails")
	}
	if a.offRoadTires {
		fmt.Println("  - Off-road tires provide good traction")
	}
	fmt.Println("  ✓ Balanced performance on/off road")
}