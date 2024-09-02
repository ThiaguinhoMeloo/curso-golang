package controlstructures

import (
	"fmt"
)

func ControlStructures() {
	number := 10

	if number > 15 {
		fmt.Println("Maior que 15, número: ", number)
	} else {
		fmt.Println("Menor ou igual a 15! número: ", number)
	}

	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Número é maior que zero! número: ", anotherNumber)
	} else if number < -10 {
		fmt.Println("Número é menor que - 10! número: ", anotherNumber)
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
