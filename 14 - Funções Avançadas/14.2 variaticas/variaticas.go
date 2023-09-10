package main

//variaticas são funções que recebem uma quantidade variável de parâmetros
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

//só pode haver um parâmetro variático por função e ele deve ser o último parâmetro
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9)

	println(totalDaSoma)

	escrever("Olá mundo", 1, 2, 3, 4, 5, 6, 7, 8, 9)
}
