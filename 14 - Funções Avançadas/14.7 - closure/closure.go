package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

//função closure é uma função que referencia variáveis de fora do seu corpo
func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
