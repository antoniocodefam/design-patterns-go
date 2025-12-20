package main

type CityMotorcycleFactory struct {
}

func (cmf *CityMotorcycleFactory) createMotorcycle() IMotorcycle {
	return &Scooter{
		Motorcycle: Motorcycle{
			design: "Step-through, foot platform",
			power: 50.0,
		},
	}
}

func (cmf *CityMotorcycleFactory) createTire() ITire {
	return  &CityTire{
		Tire: Tire{
			size: 14,
			tireType: City,
		},
	}
}