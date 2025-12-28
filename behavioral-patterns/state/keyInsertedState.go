package main

import "fmt"

type KeyInsertedState struct {
	motorcycle *motorcycle
}

func (s *KeyInsertedState) insertKey() error {
	return fmt.Errorf("key is already inserted")
}

func (s *KeyInsertedState) turnKey() error {
	fmt.Println("Key turned - electronics powered on")
	fmt.Println("Dashboard lights up")
	return nil
}

func (s *KeyInsertedState) pressStartButton() error {
	fmt.Println("Pull clutch and press start button")
	fmt.Println("Engine cranking...")
	fmt.Println("Engine started! Vroom vroom!")
	s.motorcycle.setState(s.motorcycle.engineRunningState)
	return nil
}

func (s *KeyInsertedState) releaseClutch() error {
	return fmt.Errorf("cannot release clutch - engine not running")
}

func (s *KeyInsertedState) getStateString() string {
	return "Key Inserted (Ready to Start)"
}