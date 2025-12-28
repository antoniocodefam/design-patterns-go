package main

type SportMode struct{}

func (s *SportMode) getThrottleResponse(input int) int {
	// Aggressive throttle response - 90% of input
	return (input * 9) / 10
}

func (s *SportMode) getTractionControl() int {
	return 3 // Moderate intervention
}

func (s *SportMode) getABSLevel() int {
	return 2 // Reduced ABS for better feel
}

func (s *SportMode) getPowerOutput() int {
	return 100 // Full power
}

func (s *SportMode) getName() string {
	return "Sport"
}