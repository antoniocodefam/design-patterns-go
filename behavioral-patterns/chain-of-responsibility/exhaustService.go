package main

import "fmt"

type ExhaustService struct{
	ServiceStep
}

func (e *ExhaustService) proceed(vehicle *VehicleService) {
	if vehicle.exhaustCleaned {
		fmt.Println("Ehaust is alreay clean")
	} else {
		fmt.Println("Cleaning exhaust")
		vehicle.exhaustCleaned = true
	} 

	if e.next == nil {
		fmt.Println("Nothing left to do, service is done!")
	} else {
		e.next.proceed(vehicle)
	}
}