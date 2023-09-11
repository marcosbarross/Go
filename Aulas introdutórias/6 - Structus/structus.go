package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func main() {
	var u usuario
	u.nome = "João"
	u.idade = 30

	usuario2 := usuario{"Maria", 25}

	//para criar um struct sem inicializar todos os campos, é necessário informar o nome do campo
	usuario3 := usuario{idade: 40}

	fmt.Println(u)
	fmt.Println(usuario2)
	fmt.Println(usuario3)
}
