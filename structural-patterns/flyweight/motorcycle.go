package main

type motorcycle struct{
	power int
	manufactureYear int
	model IMotorcycleModel
}

func newMotocycle(model IMotorcycleModel, power, manufactureYear int) IMotorcycle{
	return &motorcycle{
		model: model,
		power: power,
		manufactureYear: manufactureYear,
	}
}

func (m *motorcycle) getPower() int{
	return m.power
}

func (m *motorcycle) getManufactureYear() int{
	return m.manufactureYear
}

func (m *motorcycle) getModel() IMotorcycleModel{
	return m.model
}