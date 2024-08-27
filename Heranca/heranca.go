package heranca

import (
	"fmt"
)

type people struct {
	name    string
	surname string
	age     int32
	height  int32
}

type student struct {
	people
	course  string
	faculty string
}

func NotHeritage() {
	fmt.Println("teste")

	p1 := people{
		name:    "João",
		surname: "Pedro",
		age:     20,
		height:  178,
	}

	s1 := student{
		people:  p1,
		course:  "Ciência da Computação",
		faculty: "UNICAMP",
	}
	fmt.Println(s1)
}

// import (
// 	"fmt"
// 	"module/heranca"
// )

// func main() {
// 	fmt.Println("Arquivo structs")
// 	heranca.NotHeritage()
// }
