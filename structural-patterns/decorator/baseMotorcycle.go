package main

type baseMotorcycle struct{
	price       float32
	description string
}

func newBaseMotorcycle(price float32, description string) IMotorcycle {
	return &baseMotorcycle{
		price: price,
		description: description,
	}
}

func (m *baseMotorcycle) getDescription() string{
	return m.description
}

func (m *baseMotorcycle) getPrice() float32{
	return m.price
}