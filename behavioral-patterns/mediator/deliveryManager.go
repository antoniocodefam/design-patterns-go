package main

import "fmt"

const MAX_TRUCK = 2

type DeliveryManager struct{
	trucks []DeliveryTruck
	waitingTrucks []DeliveryTruck
}

func (m *DeliveryManager) canTruckEnter(truck DeliveryTruck) bool{
	if len(m.trucks) < MAX_TRUCK {
		m.trucks = append(m.trucks, truck)
		return true
	}

	fmt.Println("Max trucks allowed limit reached")
	m.waitingTrucks = append(m.waitingTrucks, truck)
	return false
}

func (m *DeliveryManager) truckDeparted(truck DeliveryTruck){
	for i, t := range m.trucks {
		if t.id == truck.id {
			fmt.Printf("Truck %d departed, remaining %d trucks, waiting %d trucks\n", truck.id, len(m.trucks)-1,len(m.waitingTrucks))
			m.trucks =  append(m.trucks[0:i], m.trucks[i+1:]...)
			
			if len(m.waitingTrucks) > 0 {
				nextTruck := m.waitingTrucks[0]
				fmt.Printf("Allowing truck %d to arrive\n", nextTruck.id)
				m.waitingTrucks = m.waitingTrucks[1:]
                nextTruck.permitArrival()
			}
			return
		}
	}	

	fmt.Printf("Truck %d not found\n",truck.id)
}