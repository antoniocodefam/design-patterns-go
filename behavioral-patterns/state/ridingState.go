package main

import "fmt"

type RidingState struct {
	motorcycle *motorcycle
}

func (s *RidingState) insertKey() error {
	return fmt.Errorf("cannot insert key while riding")
}

func (s *RidingState) turnKey() error {
	return fmt.Errorf("cannot turn off engine while riding - stop first!")
}

func (s *RidingState) pressStartButton() error {
	return fmt.Errorf("engine is already running and motorcycle is in motion")
}

func (s *RidingState) releaseClutch() error {
	fmt.Println("Shifting gears...")
	return nil
}

func (s *RidingState) getStateString() string {
	return "Riding (In Motion)"
}
