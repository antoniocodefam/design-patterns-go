package main

type luggageMotorcycleDecorator struct {
	motorcycleDecorator
}

func (m *luggageMotorcycleDecorator) getDescription() string{
	return m.wrapped.getDescription() + ", Top box + Side cases"
}

func (m *luggageMotorcycleDecorator) getPrice() float32 {
	return m.wrapped.getPrice() + 730.00
}

func newLuggageMotorcycleDecorator(wrapped IMotorcycle) *luggageMotorcycleDecorator{
	return &luggageMotorcycleDecorator{
		motorcycleDecorator: motorcycleDecorator{
			wrapped: wrapped,
		},
	}
}