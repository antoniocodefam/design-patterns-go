package main

const DELIMITER = "|"

type motorcycleModelFactory struct{
	motorcycleModels map[string]IMotorcycleModel
}

var factory = &motorcycleModelFactory{}

func getMotorcycleModelFactory() *motorcycleModelFactory{
	return factory
}

func (f *motorcycleModelFactory) getModel(brand, model string) IMotorcycleModel {
	if f.motorcycleModels == nil {
		f.motorcycleModels = make(map[string]IMotorcycleModel)
	}

	key := brand + DELIMITER + model
	m, ok := f.motorcycleModels[key]

	if !ok {
		m = newMotorcycleModel(brand, model)
		f.motorcycleModels[key] = m
	}

	return m
}
