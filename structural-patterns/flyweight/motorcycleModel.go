package main

type motorcycleModel struct{
	brand string
	model string
}

func newMotorcycleModel(brand, model string) IMotorcycleModel{
	return &motorcycleModel{
		brand: brand,
		model: model,
	}
}

func (m *motorcycleModel) getBrand() string{
	return m.brand
}

func (m *motorcycleModel) getModel() string{
	return m.model
}