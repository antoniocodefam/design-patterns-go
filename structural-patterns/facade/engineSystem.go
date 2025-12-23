package main

import "fmt"

type EngineSystem struct{}

func (e *EngineSystem) checkEngine() string {
	fmt.Println("  Checking engine oil level...")
	fmt.Println("  Checking air filter...")
	fmt.Println("  Checking spark plugs...")
	return "Engine: OK"
}

