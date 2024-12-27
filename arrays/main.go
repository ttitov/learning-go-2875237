package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Yellow"
	colors[2] = "Green"

	fmt.Println(colors)
	fmt.Println(colors[0])
	fmt.Println(colors[1])

	var numbers = [5]int{3, 5, 6, 12, 44}
	fmt.Println(numbers)
	fmt.Println("The amount of numbers", len(numbers))

}
