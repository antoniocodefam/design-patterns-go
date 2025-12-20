package main

type Director struct{
  builder IVehicleBuilder
}

func (d *Director) build() {
	d.builder.setTires()
	d.builder.setSeats()
	d.builder.setSteeringWheel()
}

func newDirector(builder IVehicleBuilder) *Director {
	return &Director{
		builder: builder,
	}
}