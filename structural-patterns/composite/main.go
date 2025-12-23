package main

import (
	"fmt"
	"strings"
)

func main() {
	front := newAssembly("Front Assembly")
	front.add(newComponent("Front Wheel", 180.20))
	front.add(newComponent("Front Brake", 95.50))
	front.add(newComponent("Handlebar", 120.00))

	rear := newAssembly("Rear Assembly")
	rear.add(newComponent("Rear Wheel", 170.99))
	rear.add(newComponent("Rear Brake", 88.40))
	rear.add(newComponent("Chain & Sprocket", 75.00))

	powertrain := newAssembly("Powertrain")
	powertrain.add(newComponent("Engine 650cc", 3200.00))
	powertrain.add(newComponent("Exhaust", 420.00))
	powertrain.add(newComponent("Fuel System", 180.00))

	chassis := newAssembly("Motorcycle Chassis")
	chassis.add(front)
	chassis.add(rear)
	chassis.add(powertrain)
	chassis.add(newComponent("Frame", 800.00))
	chassis.add(newComponent("Seat", 90.00))
	chassis.add(newComponent("Lights", 140.00))

	fmt.Printf("\n%-40s %10.2f\n", "Total price", chassis.getPrice())

	chassis.remove("Rear Assembly")

	fmt.Println(strings.Repeat("-",50))
    fmt.Printf("\n%-40s %10.2f\n", "Total price without rear assembly", chassis.getPrice())
}