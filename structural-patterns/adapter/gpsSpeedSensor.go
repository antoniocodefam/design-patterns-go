package main

type GpsSpeedSensor struct{
	// Km/h speed
	speed float32
}

func (s *GpsSpeedSensor) getSpeed() float32{
	return s.speed
}