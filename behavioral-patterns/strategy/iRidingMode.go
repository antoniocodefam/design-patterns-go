package main

type IRidingMode interface{
	getThrottleResponse(input int) int 
	getABSLevel() int
	getPowerOutput() int
	getName() string
	getTractionControl() int
}