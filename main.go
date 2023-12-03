package main

import (
	"fmt"
	"math/rand"
	"time"
)

var number int

func main() {
	for i := 0; i <= 1000; i++ {
		if i%100 == 0 {
			rand.Seed(time.Now().UnixNano())
			number = rand.Intn(100) // Set number to a random value and limit it for readability
			fmt.Printf("Index %d, New Number: %d\n", i, number)
		}
	}
	// Calculate the square of the number
	fmt.Println("The square of the number is:", number*number)
}

func add(a, b int) int {
	return a + b
}
