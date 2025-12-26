package main

type MotorcycleIterator struct{
	motorcycles []*Motorcycle
	index int
}

func (m *MotorcycleIterator) getNext() *Motorcycle{
	if m.hasMore() {
		motorcycle := m.motorcycles[m.index]
		m.index++
		return motorcycle
	}

	return nil
}

func (m *MotorcycleIterator) hasMore() bool{
	if m.index < len(m.motorcycles) {
		return  true
	}

	return false
}