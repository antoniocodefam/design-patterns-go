package main

import "fmt"

type ElectricalSystem struct{}

func (e *ElectricalSystem) checkElectrical() string {
	fmt.Println("  Checking battery voltage...")
	fmt.Println("  Testing lights...")
	fmt.Println("  Checking wiring...")
	return "Electrical: OK"
}

