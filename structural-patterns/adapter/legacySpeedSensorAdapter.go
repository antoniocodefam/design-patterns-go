package main

type LegacySpeedSensorAdapter struct{
	legacySensor *LegacySpeedSensor
}

const MPH_TO_KMH_FACTOR = 1.60934
func (adapter *LegacySpeedSensorAdapter) getSpeed() float32{
   return float32(adapter.legacySensor.mphSpeed) / MPH_TO_KMH_FACTOR
}