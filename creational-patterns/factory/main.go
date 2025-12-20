package main

import "fmt"

func main() {
	var kia, _ = getCar("Kia")
	var mg, _ = getCar("MG")

	var cars = []ICar{kia,mg}

	for i,car := range cars {
        fmt.Printf("%d.", i+1)
		car.printDetails()
	}
}