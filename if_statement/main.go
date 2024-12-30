package main

import (
	"fmt"
)

func main() {
	var result string
	theAnswer := 42

	if theAnswer < 0 {
		result = "The answer less than zero"
	} else if theAnswer == 0 {
		result = "The answer is zero"
	} else {
		result = "The answer greater than zero"
	}

	fmt.Println(result)

	if x := -42; x < 0 {
		result = "The answer less than zero"
	} else if x == 0 {
		result = "The answer is zero"
	} else {
		result = "The answer greater than zero"
	}

	fmt.Println(result)

}
