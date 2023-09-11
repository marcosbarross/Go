package main

import (
	"fmt"
)

func main() {
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	var array2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	//Ao usar os 3 pontos o compilador conta quantos elementos tem no array
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array3)

	//O slice é uma fatia de um array
	//Arrays e slices são de tipos diferentes
	slice := []int{1, 2, 3, 4, 5}

	//O append recebe como parametro um slice e um elemento a ser adicionado na última posição
	slice = append(slice, 6)

	fmt.Println(slice)

	//O pode pegar uma fatia do array, o primeiro indice é inclusivo e o segundo indice é exclusivo
	slice2 := array2[1:3]
	fmt.Println(slice2)

	//Arrays internos
	fmt.Println("Arrays internos: ")
	slice3 := make([]float32, 10, 15) //tipo, tamanho e capacidade
	fmt.Println(slice3)
	fmt.Println("Tamanho do slice: ", len(slice3))
	fmt.Println("Capacidade do slice: ", cap(slice3))

	//Se estourar o tamanho do slice, ele dobra a capacidade
}
