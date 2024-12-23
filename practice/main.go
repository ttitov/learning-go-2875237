package main

import (
	"bufio"
	"fmt"
	"os"
)

const aConst int = 64

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You enter: ", input)
}
