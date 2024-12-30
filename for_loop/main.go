package main

import (
	"fmt"
)

func main() {

	colors := []string{"red", "blue", "white"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, v := range colors {
		fmt.Println(v)
	}

	//fmt.Println("some text")
}
