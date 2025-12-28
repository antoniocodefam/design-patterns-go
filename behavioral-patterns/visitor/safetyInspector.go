package main

import "fmt"

type SafetyInspector struct{}

func (i *SafetyInspector) inspectSportBike(s *SportBike){
	fmt.Printf("Safety inspection: %s\n", s.name)
	fmt.Println("  - Checking brake system (critical for high speeds)")
	fmt.Println("  - Inspecting tire condition")
	if s.aerodynamics {
		fmt.Println("  - Verifying aerodynamic fairings are secure")
	}
	fmt.Println("  - Testing traction control system")
	fmt.Println("  ✓ Sport bike passes safety inspection")
}

func (i *SafetyInspector) inspectCruiser(c *Cruiser){
	fmt.Printf("Safety inspection: %s\n", c.name)
	fmt.Println("  - Checking brake system")
	fmt.Println("  - Inspecting lights and signals")
	if c.saddlebags {
		fmt.Println("  - Ensuring saddlebag mounts are secure")
	}
	fmt.Println("  - Testing horn and mirrors")
	fmt.Println("  ✓ Cruiser passes safety inspection")
}

func (i *SafetyInspector) inspectAdventureBike(a *AdventureBike){
	fmt.Printf("Safety inspection: %s\n", a.name)
	fmt.Println("  - Checking brake system")
	if a.offRoadTires {
		fmt.Println("  - Inspecting off-road tire tread depth")
	}
	if a.luggage {
		fmt.Println("  - Verifying luggage attachment security")
	}
	fmt.Println("  - Testing ABS and traction control")
	fmt.Println("  ✓ Adventure bike passes safety inspection")
}