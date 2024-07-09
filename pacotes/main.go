package main

import (
	"fmt"
	"oo/pacotes/car"
)

func main() {
	car := car.Car{"Gol", "Yellow"}
	fmt.Println(car.Start())
}
