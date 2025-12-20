package main

type MotorcycleBuilder struct{
    vehicle Vehicle
}

func (b *MotorcycleBuilder) setSeats() {
    b.vehicle.seats = 1
}

func (b *MotorcycleBuilder) setTires() {
	b.vehicle.tires = 2
}

func (b *MotorcycleBuilder) setSteeringWheel() {
	b.vehicle.steeringWheel = "Handlebar"
}

func (b *MotorcycleBuilder) reset() {
	b.vehicle = Vehicle{}
}

func (b *MotorcycleBuilder) getVehicle() Vehicle {
	vehicle := b.vehicle
    b.reset()

	return  vehicle
}

func newMotorcycleBuilder() *MotorcycleBuilder {
	return &MotorcycleBuilder{}
}