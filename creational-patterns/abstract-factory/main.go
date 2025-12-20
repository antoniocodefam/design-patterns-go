package main

import (
	"fmt"
	"strings"
)

func main() {
	dirtFactory, _ := GetVehicleFactory("dirt")
	cityFactory, _ := GetVehicleFactory("city")

	PrintDetails(dirtFactory.createMotorcycle(),dirtFactory.createTire())
	fmt.Println(strings.Repeat("-", 70))
    PrintDetails(cityFactory.createMotorcycle(),cityFactory.createTire())    
}

func PrintDetails(motorcycle IMotorcycle, tire ITire) {
	fmt.Printf("Motorcycle// Power: %.2f - Design: %s",motorcycle.getPower(), motorcycle.getDesign())
	fmt.Println()
	fmt.Printf("Tire// Size: %d - Type: %s", tire.getSize(),tire.getType())
	fmt.Println()
}