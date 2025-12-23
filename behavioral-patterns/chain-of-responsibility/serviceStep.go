package main

import "fmt"

type ServiceStep struct {
	next IServiceStep
}

func (s *ServiceStep) setNext(next IServiceStep) {
	s.next = next
}

func (s *ServiceStep) proceed(vehicle *VehicleService) {
	if s.next == nil {
		fmt.Println("Nothing left to do, service is done!")
	}
	s.next.proceed(vehicle)
}