package main

import "fmt"

type OilChange struct{
	ServiceStep
}

func (o *OilChange) proceed(vehicle *VehicleService) {
	if vehicle.oilChanged {
		fmt.Println("Oil is already changed")
	} else {
		fmt.Println("Changing oil")
		vehicle.oilChanged = true
	} 

	if o.next == nil {
		fmt.Println("Nothing left to do, service is done!")
	} else {
		o.next.proceed(vehicle)
	}
}