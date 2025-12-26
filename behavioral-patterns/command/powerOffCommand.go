package main

type PowerOffCommand struct{
	vehicle IVehicle
}

func (c *PowerOffCommand) execute() {
	c.vehicle.off()
}