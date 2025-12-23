package main

import "fmt"

type ServiceReport struct{}

func (r *ServiceReport) generateReport(id string, engine, brakes, tires, electrical string) {
	fmt.Println("\n=== SERVICE REPORT ===")
	fmt.Printf("Motorcycle ID: %s\n", id)
	fmt.Printf("  %s\n", engine)
	fmt.Printf("  %s\n", brakes)
	fmt.Printf("  %s\n", tires)
	fmt.Printf("  %s\n", electrical)
	fmt.Println("====================")
}

