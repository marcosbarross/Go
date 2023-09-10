package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//O if pode ter uma inicialização de variável que só vai existir dentro do escopo do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	}

}
