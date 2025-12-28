package main

import "fmt"

func main() {
	motorcycle := newMotorcycle()

	fmt.Println("Current state: ", motorcycle.getCurrentState())
	motorcycle.insertKey()

	fmt.Println("Current state: ", motorcycle.getCurrentState())
	motorcycle.turnKey()

	fmt.Println("Current state: ", motorcycle.getCurrentState())
	motorcycle.pressStartButton()

	fmt.Println("Current state: ", motorcycle.getCurrentState())
	motorcycle.releaseClutch()

	fmt.Println("Attempting invalid actions")
	fmt.Println("Current state: ", motorcycle.getCurrentState())
	err := motorcycle.insertKey()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	err = motorcycle.pressStartButton()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}