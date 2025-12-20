package main

type IMotorcycle interface {
	getDesign() string
	setDesign(design string)
	getPower() float32
	setPower(power float32)
}
