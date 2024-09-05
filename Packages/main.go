package main

import (
	"fmt"
	"module/advancedfunctions/namedreturn"
)

func main() {
	soma, subtracao := namedreturn.NamedReturn(10, 20)
	fmt.Println(soma, subtracao)
}
