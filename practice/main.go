package main

import (
	"fmt"
)

func main() {
	i1, i2, i3 := 1, 3, 45

	intSum := i1 + i2 + i3
	fmt.Println("Integers sum is ", intSum)

	f1, f2, f3 := 23.5, 55.1, 6.3

	floatSum := f1 + f2 + f3
	fmt.Println("Float sum is ", floatSum)

}
