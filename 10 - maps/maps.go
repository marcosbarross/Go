package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//Maps são como dicionários em Python e os tipos precisam ser sempre os mesmos, o primeiro parametro é a chave e o segundo o valor
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	//aninhamento de maps
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
}
