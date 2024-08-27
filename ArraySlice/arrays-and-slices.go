package arrayslice

import (
	"fmt"
	"reflect"
)

func ArraySlice() {
	fmt.Println("Arrays e Slices")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
	// Array -> OBS: no Arrya limitamos o tamanho, não podemos deixar de colocar um limite dentro de um array.
	fmt.Println("Array:")
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
	// Slice -> OBS: Não precisamos nos preocupar com o limite de tamanho na hora de definir / criar um slice. Slice é mais utilizado que o array.
	fmt.Println("Slices:")
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// mostrando o tipo do slice e array
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	// referenciando um array
	slice2 := array2[1:3] // [1:3] -> 1 é um indice inclusivo, 3 é um indice é exclusivo
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
	//Arrays Internos
	fmt.Println("Arrays Internos:")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length -> ver o tamnho
	fmt.Println(cap(slice3)) // cap -> ver a capacidade maxima

	slice4 := make([]float32, 5)
	slice4 = append(slice4, 6)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
}
