package main

func main() {
	var variavel1 string = "Variável 1"

	//variável com inferência de tipo
	variavel2 := "Variável 2"

	//opção de declaração de variáveis
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	println(variavel1, variavel2)

	println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	println(variavel5, variavel6)

	//troca de valores entre variáveis
	variavel5, variavel6 = variavel6, variavel5
	println(variavel5, variavel6)

	//declaração de constante, não pode ser alterada
	const constante1 string = "Constante 1"
	println(constante1)
}
