package main

type Tire struct {
	size     int
	tireType TireType
}

func (t *Tire) getSize() int {
	return t.size
}

func (t *Tire) setSize(size int) {
	t.size = size
}

func (t *Tire) getType() TireType {
	return t.tireType
}

func (t *Tire) setType(tireType TireType) {
	t.tireType = tireType
}