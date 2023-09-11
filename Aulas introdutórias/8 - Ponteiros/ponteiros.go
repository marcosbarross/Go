package main

import "fmt"

func main() {
	var variavale1 int = 10
	var variavel2 int = variavale1

	fmt.Println(variavale1, variavel2)

	variavale1++
	fmt.Println(variavale1, variavel2)

	//ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	//ponteiro recebe o endereço de memória da variavel3 ao usar o &
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) //se usar o * antes do ponteiro, ele vai imprimir o valor da variavel3

	//o ponteiro é usado para passar informações para uma função sem fazer uma cópia de uma variável em um momento específico, dessa forma garantindo que o valor da variável estará atualizado
}
