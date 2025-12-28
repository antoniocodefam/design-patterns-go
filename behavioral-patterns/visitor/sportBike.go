package main

type SportBike struct {
	name          string
	engineCC      int
	topSpeed      int
	aerodynamics  bool
	raceMode      bool
}

func (s *SportBike) inspect(i IInspector) {
	i.inspectSportBike(s)
}

func (s *SportBike) getName() string {
	return s.name
}