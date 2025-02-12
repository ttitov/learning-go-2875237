// Write your answer here, and then test your code.

package main

import "fmt"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	// sum := 0.0
	// for i := 0; i < len(cart); i++ {
	// 	sum += cart[i].price * float64(cart[i].quantity)
	// }

	var total float64 = 0
	for _, k := range cart {
		total += k.price * float64(k.quantity)
	}

	return total
}

func main() {
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)

	fmt.Println(result)
}
