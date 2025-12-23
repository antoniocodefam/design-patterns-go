package main

type IMotorcycleModel interface{
	getBrand() string
	getModel() string
}

type IMotorcycle interface{
	getPower() int
	getManufactureYear() int
	getModel() IMotorcycleModel
}