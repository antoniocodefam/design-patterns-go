package main

type RainMode struct{}

func (r *RainMode) getThrottleResponse(input int) int {
	// Very gentle throttle response - 50% of input
	return input / 2
}

func (r *RainMode) getTractionControl() int {
	return 9 // Maximum intervention (scale 0-9)
}

func (r *RainMode) getABSLevel() int {
	return 3 // Maximum ABS intervention (scale 0-3)
}

func (r *RainMode) getPowerOutput() int {
	return 50 // 50% power limit
}

func (r *RainMode) getName() string {
	return "Rain"
}