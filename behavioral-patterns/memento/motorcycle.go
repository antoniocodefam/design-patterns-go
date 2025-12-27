package main

import "fmt"

type Motorcycle struct{
	model string
	tirePressure float32
	absLevel int
	tractionControl int
}

func (m *Motorcycle) createConfig() *MotorcycleConfig{
	return &MotorcycleConfig{
		tirePressure: m.tirePressure,
		absLevel: m.absLevel,
		tractionControl: m.tractionControl,
	}
}

func (m *Motorcycle) setConfig(config MotorcycleConfig) {
	m.absLevel = config.absLevel
	m.tirePressure = config.tirePressure
	m.tractionControl = config.tractionControl
}

func (m *Motorcycle) setTirePressure(tirePressure float32) {
	m.tirePressure = tirePressure
}
func (m *Motorcycle) setAbsLevel(level int) {
	m.absLevel = level
}
func (m *Motorcycle) setTractionControl(tractionControl int) {
	m.tractionControl = tractionControl
}

func (m *Motorcycle) displayConfig() {
	fmt.Printf("Tire Pressure: %.2f, ABS Level: %d, Traction Control; %d\n", m.tirePressure, m.absLevel, m.tractionControl)
}