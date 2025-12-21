package main

type LegacySpeedSensor struct{
	mphSpeed int
}

func (s *LegacySpeedSensor) readSpeed() int{
	return  s.mphSpeed
}