package main

import "fmt"

type Vehicle interface {
	start() string
}

type Car struct {
	name string
}

type Motocycle struct {
	name string
}

func startVehicle(vehicle Vehicle) string {
	return vehicle.start()
}

func (car Car) start() string {
	return "The car" + car.name + "has been started"
}

func (motocycle Motocycle) start() string {
	return "The motorcycle " + motocycle.name + " has been started"
}

func main() {
	car := Car{"Gol"}
	motocycle := Motocycle{"XPTO"}
	fmt.Println(startVehicle(car))
	fmt.Println(startVehicle(motocycle))
}
