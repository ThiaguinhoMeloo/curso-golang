package main

import (
	"fmt"
	"module/advancedfunctions/recursive"
)

func main() {
	position := uint(12)
	for i := uint(0); i <= position; i++ {
		fmt.Println(recursive.Fibonacci(position))
	}
}
