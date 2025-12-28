package main

import "fmt"

type EngineRunningState struct {
	motorcycle *motorcycle
}

func (s *EngineRunningState) insertKey() error {
	return fmt.Errorf("key is already in ignition and engine is running")
}

func (s *EngineRunningState) turnKey() error {
	fmt.Println("Turning off engine")
	s.motorcycle.setState(s.motorcycle.keyInsertedState)
	return nil
}

func (s *EngineRunningState) pressStartButton() error {
	return fmt.Errorf("engine is already running")
}

func (s *EngineRunningState) releaseClutch() error {
	fmt.Println("Clutch released - engaging gear")
	fmt.Println("Motorcycle is now moving!")
	s.motorcycle.setState(s.motorcycle.ridingState)
	return nil
}

func (s *EngineRunningState) getStateString() string {
	return "Engine Running (Stationary)"
}