package main

type motorcycle struct {
	state IMotorcycleState

	powerOffState IMotorcycleState
	keyInsertedState IMotorcycleState
	engineRunningState IMotorcycleState
	ridingState IMotorcycleState
}

func newMotorcycle() *motorcycle {
	m := &motorcycle{}
	m.state = &PowerOffState{ motorcycle: m }

	m.powerOffState = m.state
	m.engineRunningState = &EngineRunningState{motorcycle: m}
	m.keyInsertedState = &KeyInsertedState{motorcycle: m}
	m.ridingState = &RidingState{motorcycle: m}
	
	return m
}

func (m *motorcycle) setState(state IMotorcycleState)  {
	m.state = state
}

func (m *motorcycle) insertKey() error {
	return m.state.insertKey()
}
func (m *motorcycle) turnKey() error {
	return m.state.turnKey()
}
func (m *motorcycle) pressStartButton() error {
	return m.state.pressStartButton()
}
func (m *motorcycle) releaseClutch() error {
	return m.state.releaseClutch()
}
func (m *motorcycle) getCurrentState() string {
	return m.state.getStateString()
}


