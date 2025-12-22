package main

type heatedGripsMotorcycleDecorator struct {
	motorcycleDecorator
}

func (m *heatedGripsMotorcycleDecorator) getDescription() string{
	return m.wrapped.getDescription() + ", Heated Grips"
}

func (m *heatedGripsMotorcycleDecorator) getPrice() float32 {
	return m.wrapped.getPrice() + 39.99
}

func newHeatedGripsMotorcycleDecorator(wrapped IMotorcycle) *heatedGripsMotorcycleDecorator{
	return &heatedGripsMotorcycleDecorator{
		motorcycleDecorator: motorcycleDecorator{
			wrapped: wrapped,
		},
	}
}