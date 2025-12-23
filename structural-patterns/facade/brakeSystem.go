package main

import "fmt"

type BrakeSystem struct{}

func (b *BrakeSystem) checkBrakes() string {
	fmt.Println("  Checking brake pads...")
	fmt.Println("  Checking brake fluid...")
	fmt.Println("  Testing brake responsiveness...")
	return "Brakes: OK"
}

