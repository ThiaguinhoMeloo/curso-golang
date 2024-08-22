package structs

import (
	"fmt"
)

type user struct {
	name    string
	age     int32
	address address
}

type address struct {
	logradouro string
	number     int32
}

func Struct() {
	var u user
	u.name = "Davi"
	u.age = 21
	fmt.Println(u)

	exampleAddress := address{
		logradouro: "Av.Palestina",
		number:     7000,
	}

	usuario2 := user{
		name:    "Davi",
		age:     22,
		address: exampleAddress,
	}
	fmt.Println(usuario2)
}

// import (
// 	"fmt"
// 	"module/structs"
// )

// func main() {
// 	fmt.Println("Arquivo structs")
// 	structs.Struct()
// }
