package main

import "fmt"

func main() {

	// slice := make([]int, 2, 5)
	// slice = append(slice, 1, 2, 3)
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	sliceString := []string{
		"Helo",
		"World",
		"Much",
		"Better",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0])
	fmt.Println(sliceString[:2])
	fmt.Println(sliceString[1:2])
	fmt.Println(sliceString[2:4])
}
