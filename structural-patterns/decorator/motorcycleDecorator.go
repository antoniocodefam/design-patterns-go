package main

type motorcycleDecorator struct{
	wrapped IMotorcycle
}

func (md *motorcycleDecorator) getDescription() string{
	return md.wrapped.getDescription()
}

func (md *motorcycleDecorator) getPrice() float32{
	return md.wrapped.getPrice()
}