package main

type motorcycle struct{
	ridingMode IRidingMode
	maxPower int
	model string
}

func newMotorcycle(model string, maxPower int) *motorcycle {
	return &motorcycle{
		model: model,
		maxPower: maxPower,
		ridingMode: &StreetMode{},
	}
}

func (m *motorcycle) setRidingMode(ridingMode IRidingMode) {
	m.ridingMode = ridingMode
}

func (m *motorcycle) getThrottleResponse() int {
	return m.ridingMode.getThrottleResponse(m.maxPower)
}

func (m *motorcycle) getABSLevel() int {
	return m.ridingMode.getABSLevel()
}

func (m *motorcycle) getPowerOutput() int {
	return m.ridingMode.getPowerOutput()
}

func (m *motorcycle) getTractionControl() int {
	return m.ridingMode.getTractionControl()
}

func (m *motorcycle) getRidingModeName() string {
	return m.ridingMode.getName()
}