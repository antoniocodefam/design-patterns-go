package main

import "fmt"

func getCar(brand string) (ICar, error) {
	if brand == "Kia" {
		return newKia(), nil
	}

	if brand == "MG" {
		return newMg(), nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}