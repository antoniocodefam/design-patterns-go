package main

type DirtMotorcycleFactory struct{

}

func (dmf *DirtMotorcycleFactory) createMotorcycle() IMotorcycle {
	return  &DirtBike{
		Motorcycle: Motorcycle{
			power: 300,
			design: "Cradle, high ground clearance",
		},
	}
}

func (dmf *DirtMotorcycleFactory) createTire() ITire {
	return &OffRoadTire{
		Tire: Tire{
			size: 17,
			tireType: OffRoad,
		},
	}
}