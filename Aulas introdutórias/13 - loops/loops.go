package main

import (
	"fmt"
	//"time"
)

func main() {
	i := 0

	for i < 1 {
		i++
		fmt.Println("Incrementando i")
		//time.Sleep(time.Second)
	}

	for j := 0; j < 1; j++ {
		fmt.Println("Incrementando j", j)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		//para retornar a letra é preciso converter para string, caso contrário será retornado o valor da tabela ASCII
		fmt.Println("com conversão: ", indice, string(letra))
		//exemplo sem conversão
		fmt.Println("sem conversão: ", indice, letra)
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//for infinito
	//for {
	//	fmt.Println("Executando infinitamente")
	//	time.Sleep(time.Second)
	//}
}
