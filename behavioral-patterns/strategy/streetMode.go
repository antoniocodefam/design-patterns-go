package main

type StreetMode struct{}

func (s *StreetMode) getThrottleResponse(input int) int {
	// Moderate throttle response - 70% of input
	return (input * 7) / 10
}

func (s *StreetMode) getTractionControl() int {
	return 7 // High intervention
}

func (s *StreetMode) getABSLevel() int {
	return 3 // Full ABS
}

func (s *StreetMode) getPowerOutput() int {
	return 80 // 80% power limit
}

func (s *StreetMode) getName() string {
	return "Street"
}