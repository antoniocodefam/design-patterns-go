package main

type absMotorcycleDecorator struct {
	motorcycleDecorator
}

func (m *absMotorcycleDecorator) getDescription() string{
	return m.wrapped.getDescription() + ", ABS"
}

func (m *absMotorcycleDecorator) getPrice() float32 {
	return m.wrapped.getPrice() + 210.99
}

func newAbsMotorcycleDecorator(wrapped IMotorcycle) IMotorcycle {
	return &absMotorcycleDecorator{
		motorcycleDecorator: motorcycleDecorator{
			wrapped: wrapped,
		},
	}
}