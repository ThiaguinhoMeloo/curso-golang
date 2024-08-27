package pointers

import (
	"fmt"
)

func Pointer() {
	fmt.Println("Ponteiros")

	var variable1 int = 10
	var variable2 int = variable1
	fmt.Println(variable1, variable2)

	variable1++
	fmt.Println(variable1, variable2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	// Quando criamos um ponteiro, não estamos jogando o valor da variavel dentro,
	// vamos estar jogando o endereço de memoria onde essa memoria esta salva.
	// O valor zero de um ponteiro é NIL.

	var variable3 int
	var pointer *int

	variable3 = 100
	pointer = &variable3
	fmt.Println(variable3, pointer)

	variable3 = 150
	fmt.Println(variable3, *pointer) // processo de desreferenciação

}

// import (
// 	"module/pointers"
// )

// func main() {
// 	pointers.Pointer()
// }
