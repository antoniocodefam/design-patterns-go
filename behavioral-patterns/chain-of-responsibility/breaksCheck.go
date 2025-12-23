package main

import "fmt"

type BreaksChange struct{
	ServiceStep
}

func (b *BreaksChange) proceed(vehicle *VehicleService) {
	if vehicle.brakesChecked {
		fmt.Println("Breaks are okay")
	} else {
		fmt.Println("Fixing breaks")
		vehicle.brakesChecked = true
	} 

	if b.next == nil {
		fmt.Println("Nothing left to do, service is done!")
	} else {
		b.next.proceed(vehicle)
	}
}