// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

func main() {

	value1 := "asasd"
	value2 := "-110.5"
	result := calculate(value1, value2)

	fmt.Println("The result is ", result)
}

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.
	// Convert the first string to a float64
	n1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be float")
	}

	// Convert the second string to a float64
	n2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be float")
	}

	// Calculate and return the result
	result := n1 + n2

	return result
}
