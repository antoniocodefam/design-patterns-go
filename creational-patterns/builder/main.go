package main

import (
	"fmt"
	"strings"
)

func main() {
	motorcycleBuilder := newMotorcycleBuilder()
	carBuilder := newCarBuilder()

	 motorcycle := build(motorcycleBuilder)
	 car := build(carBuilder)

	 fmt.Printf("Motorcycle - ")
	printDetails(motorcycle)
    fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("Car - ")
	printDetails(car)

}

func build(builder IVehicleBuilder) Vehicle {
	director := newDirector(builder)
    director.build()
	
	return builder.getVehicle()
}

func printDetails(vehicle Vehicle) {
  fmt.Printf("Seats: %d, tires: %d, steering wheel: %s",vehicle.seats,vehicle.tires,vehicle.steeringWheel)
  fmt.Println()
}