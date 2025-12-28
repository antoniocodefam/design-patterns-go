package main

import "fmt"

type PowerOffState struct {
	motorcycle *motorcycle
}

func (s *PowerOffState) insertKey() error {
	fmt.Println("Key inserted into ignition")
	s.motorcycle.setState(s.motorcycle.keyInsertedState)

	return nil
}

func (s *PowerOffState) turnKey() error {
	return fmt.Errorf("Cannot turn key - no key in ignition")
}

func (s *PowerOffState) pressStartButton() error {
	return fmt.Errorf("Cannot press start - motorcycle is off, insert key first")
}

func (s *PowerOffState) releaseClutch() error {
	return fmt.Errorf("Cannot release clutch - motorcycle is off")
}

func (s *PowerOffState) getStateString() string {
	return "Power off"
}



