package main

func main(){
	vehicleService := &VehicleService{}

	oilChangeStep := &OilChange{}
	breaksChange := &BreaksChange{}
	exhaustService := &ExhaustService{} 

	oilChangeStep.setNext(breaksChange)
	breaksChange.setNext(exhaustService)

    oilChangeStep.proceed(vehicleService)

    oilChangedVehicle := &VehicleService{
		oilChanged: true,
	}

	oilChangeStep.proceed(oilChangedVehicle)
}