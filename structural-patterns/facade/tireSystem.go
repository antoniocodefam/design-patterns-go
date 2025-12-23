package main

import "fmt"

type TireSystem struct{}

func (t *TireSystem) checkTires() string {
	fmt.Println("  Checking tire pressure...")
	fmt.Println("  Checking tire tread depth...")
	fmt.Println("  Checking tire alignment...")
	return "Tires: OK"
}

