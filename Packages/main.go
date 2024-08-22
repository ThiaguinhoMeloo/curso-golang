package main

import (
	"fmt"
	"module/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	assistant.ToWrite()

	err := checkmail.ValidateFormat("123456789")
	if err != nil {
		fmt.Println(err)
	}
}
