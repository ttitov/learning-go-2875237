package main

import (
	"fmt"
)

func main() {
	intVal := 23
	var p *int = &intVal

	fmt.Println("Some text", *p)
	fmt.Println("Value of p is", p)
	fmt.Println("Value of memory p is", &p)
	fmt.Println("Value of memory intVal is", &intVal)
}
