package main

type Kia struct {
	Car
}

func newKia() ICar {
	return &Kia {
		Car: Car{
			brand: "Kia",
			origin: "South Korea",
		},
	}
}