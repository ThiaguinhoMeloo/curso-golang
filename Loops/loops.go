package loops

import (
	"fmt"
)

func Loops() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando I")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println("Valor final de I: ", i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementeando J: ", j)
	// 	time.Sleep(time.Second)
	// 	if j == 10 {
	// 		fmt.Println("Valor final de J: ", j)
	// 	}
	// }

	names := [3]string{"João", "Davi", "Lucas"}

	for indice, name := range names {
		fmt.Println(indice, name)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	user := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range user {
		fmt.Println(chave, valor)
	}

	// Não é possivel usar range em um strutc
	// type userStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }

	// user2 := userStruct{
	// 	nome:      "Zé",
	// 	sobrenome: "Júnior",
	// }

	// for chave, valor := range user2 {
	// 	fmt.Println(chave, valor)
	// }

}
