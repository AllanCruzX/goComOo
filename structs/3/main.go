package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	color string
}

func main() {
	car := Car{"Gol", 2017, "Yellow"}

	result, _ := json.Marshal(car)
	//fmt.Println(result)
	fmt.Println(string(result))
}
