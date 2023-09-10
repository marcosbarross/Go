package main

import "fmt"

func main() {
	//função anônima, para executar ela é necessário colocar os parênteses no final
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
