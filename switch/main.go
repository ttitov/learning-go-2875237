package main

import (
	"fmt"
	"math/rand"
)

func main() {

	dow := rand.Intn(7) + 1
	var result string
	switch dow {
	case 1:
		result = "It's Monday"
		// fallthrough
	case 2:
		result = "It's Tuesday"
	case 6:
		result = "It's Saturday"
	default:
		result = "It's some day..."
	}
	fmt.Printf("Some day %v\n", dow)

	fmt.Println(result)

}
