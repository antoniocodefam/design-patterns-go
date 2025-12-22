package main

type BaseMotorcycle struct{
	price       float32
	description string
}

func (m *BaseMotorcycle) getDescription() string{
	return m.description
}

func (m *BaseMotorcycle) getPrice() float32{
	return m.price
}