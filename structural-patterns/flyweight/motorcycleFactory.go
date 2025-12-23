package main

type motorcycleFactory struct{}

func newMotorcycleFactory() *motorcycleFactory{
	return &motorcycleFactory{}
}

func (f *motorcycleFactory) createBike(model IMotorcycleModel, power, manufactureYear int) IMotorcycle{
	return &motorcycle{
		model: model,
		power: power,
		manufactureYear: manufactureYear,
	}
}