package main

type Motorcycle struct {
	design string
	power float32
}

func (m *Motorcycle) getDesign() string {
	return m.design
}

func (m *Motorcycle) setDesign(design string) {
	m.design = design
}

func (m *Motorcycle) getPower() float32 {
	return m.power
}

func (m *Motorcycle) setPower(power float32) {
	m.power = power
}