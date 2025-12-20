package main

type CarBuilder struct {
	vehicle Vehicle
}

func (b *CarBuilder) setSeats() {
	b.vehicle.seats = 4
}

func (b *CarBuilder) setTires() {
	b.vehicle.tires = 4
}

func (b *CarBuilder) setSteeringWheel() {
	b.vehicle.steeringWheel = "Round steering wheel"
}

func (b *CarBuilder) reset() {
	b.vehicle = Vehicle{}
}

func (b *CarBuilder) getVehicle() Vehicle {
	vehicle := b.vehicle
	b.reset()

	return vehicle
}

func newCarBuilder() *CarBuilder {
	return &CarBuilder{}
}