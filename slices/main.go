package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Yellow"}
	fmt.Println(colors)
	fmt.Println(colors[0])

	colors = append(colors, "Blue")

	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 6)
	for i := 0; i < 5; i++ {
		numbers[i] = i * (i + 3)
	}
	numbers[5] = 2

	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)

}
