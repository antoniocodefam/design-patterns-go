package main

type IMotorcycleState interface {
	insertKey() error
	turnKey() error
	pressStartButton() error
	releaseClutch() error
	getStateString() string
}