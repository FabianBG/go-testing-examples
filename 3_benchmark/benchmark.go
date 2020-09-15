package main

import (
	"fmt"
	"math/rand"
)

func operation() int {
	acc := 0
	for i := 0; i < 5; i++ {
		acc += rand.Int() * rand.Int()
	}

	return acc
}

func main() {
	fmt.Println(operation())
}
