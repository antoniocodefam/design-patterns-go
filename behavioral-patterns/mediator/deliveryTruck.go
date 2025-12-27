package main

import "fmt"

type DeliveryTruck struct{
	id int
	manager IManager
}

func (t *DeliveryTruck) arrive(){
	fmt.Printf("Truck %d arrived\n",t.id)

	if t.manager.canTruckEnter(*t) {
		fmt.Printf("Truck %d entering\n",t.id)
	} else {
		fmt.Printf("Truck %d cannot enter, waiting\n", t.id)
	}
}

func (t *DeliveryTruck) depart(){
	fmt.Printf("Truck %d departing\n",t.id)
	t.manager.truckDeparted(*t)
}

func (t *DeliveryTruck) permitArrival(){
	fmt.Printf("Truck %d arrival permitted\n",t.id)
	t.arrive()
}