package main

import (
	"fmt"
	"strings"
)

func main() {
	front := &Assembly{name: "Front Assembly"}
	front.add(&Component{name: "Front Wheel", price: 180.20})
	front.add(&Component{name: "Front Brake", price: 95.50})
	front.add(&Component{name: "Handlebar", price: 120.00})

	rear := &Assembly{name: "Rear Assembly"}
	rear.add(&Component{name: "Rear Wheel", price: 170.99})
	rear.add(&Component{name: "Rear Brake", price: 88.40})
	rear.add(&Component{name: "Chain & Sprocket", price: 75.00})

	powertrain := &Assembly{name: "Powertrain"}
	powertrain.add(&Component{name: "Engine 650cc", price: 3200.00})
	powertrain.add(&Component{name: "Exhaust", price: 420.00})
	powertrain.add(&Component{name: "Fuel System", price: 180.00})

	chassis := &Assembly{name: "Motorcycle Chassis"}
	chassis.add(front)
	chassis.add(rear)
	chassis.add(powertrain)
	chassis.add(&Component{name: "Frame", price: 800.00})
	chassis.add(&Component{name: "Seat", price: 90.00})
	chassis.add(&Component{name: "Lights", price: 140.00})

	fmt.Printf("\n%-40s %10.2f\n", "Total price", chassis.getPrice())

	chassis.remove("Rear Assembly")

	fmt.Println(strings.Repeat("-",50))
    fmt.Printf("\n%-40s %10.2f\n", "Total price without rear assembly", chassis.getPrice())
}