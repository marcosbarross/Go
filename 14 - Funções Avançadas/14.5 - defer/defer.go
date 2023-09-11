package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	//ao usar o defer a função é excecutada por ultimo
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		fmt.Println()
		return true
	}
	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
}
