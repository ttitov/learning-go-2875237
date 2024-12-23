package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Printf("The variable's type is %T\n", anotherInt)

	myString := "This is my string"
	fmt.Println(myString)
	fmt.Printf("This is variable's type %T\n", myString)

	fmt.Println(aConst)

}
