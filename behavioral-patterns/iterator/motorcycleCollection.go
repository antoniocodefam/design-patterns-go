package main

type MotorcycleCollection struct{
	motorcycles []*Motorcycle
}

func (m *MotorcycleCollection) createIterator() IIterator{
	return &MotorcycleIterator{
		motorcycles: m.motorcycles,
	}
}