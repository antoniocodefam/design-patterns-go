package main

import "fmt"

type chainMaintenance struct{
	BaseMaintenance
}

func NewChainMaintenance(bikeName string) *chainMaintenance{
	c := &chainMaintenance{
		BaseMaintenance: BaseMaintenance{
			bikeName: bikeName,
		},
	}

	c.procedure = c

	return c
}

func (c *chainMaintenance) performSpecificTasks() {
	fmt.Println("Chain service: clean, check tension, lubricate")
}

func (c *chainMaintenance) runDiagnostics() {
	fmt.Println("Chain tension at 20mm, no damage detected")
}