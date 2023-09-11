package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	alunoEstaAprovado(6, 6)
}
