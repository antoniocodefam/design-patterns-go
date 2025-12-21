package main

import "fmt"

func main(){
	gpsSpeedSensor := &GpsSpeedSensor{
		speed: 99.7793,
	}
	legacySpeedSensor := &LegacySpeedSensor{
		mphSpeed: 62,
	}
	adapter := &LegacySpeedSensorAdapter{
		legacySensor: legacySpeedSensor,
	}
    
	fmt.Printf("GPS Sensor:\t")
	printSpeed(gpsSpeedSensor)
	fmt.Printf("Adapter speed:\t")
	printSpeed(adapter)

}

func printSpeed(sensor ISpeedSensor) {
	fmt.Printf("Speed: %.2f\n", sensor.getSpeed())
}