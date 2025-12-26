package main

type PowerOnCommand struct{
	vehicle IVehicle
}

func (c *PowerOnCommand) execute() {
	c.vehicle.on()
}