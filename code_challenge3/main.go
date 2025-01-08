// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"strconv"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	v1 := convertInputToValue(input1)
	v2 := convertInputToValue(input2)

	if operation == "+" {
		return addValues(v1, v2)
	} else if operation == "-" {
		return subtractValues(v1, v2)
	} else if operation == "*" {
		return multiplyValues(v1, v2)
	} else if operation == "/" {
		return divideValues(v1, v2)
	}

	panic(fmt.Sprintf("Wrong operation", operation))
}

func convertInputToValue(input string) float64 {
	if res, err := strconv.ParseFloat(input, 64); err == nil {
		return res
	} else {
		panic(fmt.Sprintf("Can't convert the string to a number", input))
	}
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	if value2 != 0 {
		return value1 / value2
	} else {
		panic("The divider is zero")
	}

}

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)
	fmt.Printf(value1 + operation + value2 + "=" + strconv.FormatFloat(result, 'f', -1, 64))
}
