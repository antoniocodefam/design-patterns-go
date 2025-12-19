package main

type MG struct {
	Car
}

func newMg() ICar {
	return &MG {
		Car: Car{
			brand: "MG",
			origin: "United Kingdom",
		},
	}
}