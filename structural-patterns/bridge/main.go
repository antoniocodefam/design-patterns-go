package main

func main(){
	electricEngine := &ElectricEngine{}
	gasEngine := &GasEngine{}

	sportBike := &SportBike{}
	cruiserBike := &CruiserBike{}

	sportBike.setEngine(electricEngine)
	sportBike.start()
	sportBike.stop()

	sportBike.setEngine(gasEngine)
	sportBike.start()
	sportBike.stop()

	cruiserBike.setEngine(electricEngine)
	cruiserBike.start()
	cruiserBike.stop()

	cruiserBike.setEngine(gasEngine)
	cruiserBike.start()
	cruiserBike.stop()
}