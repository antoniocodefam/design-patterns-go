package main

import "fmt"

func main() {
	serviceFacade := newMotorcycleServiceFacade()

	err := serviceFacade.performService("MC-2024-001")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

