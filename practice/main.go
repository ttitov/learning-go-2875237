package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const aConst int = 64

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You enter: ", input)

	fmt.Print("Enter a number: ")
	inputNumber, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(inputNumber), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of the number: ", aFloat)
	}

}
