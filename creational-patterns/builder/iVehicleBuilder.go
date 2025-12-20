package main

type IVehicleBuilder interface{
	setSeats()
	setTires()
	setSteeringWheel()
	getVehicle() Vehicle
}