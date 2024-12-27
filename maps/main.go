package main

import (
	"fmt"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	fmt.Println(states)
	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)
	states["NY"] = "New York"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
}
