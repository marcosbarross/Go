package main

import "errors"

func main() {
	//tipos de int possíveis int8, int16, int32, int64
	var numero int16 = 100
	println(numero)

	//Ao utilizar inferencia de tipo o compilador usa a arquitetura do computador
	numeroInferido := 1000000000000
	println(numeroInferido)

	//STRINGS - precisa usar aspas duplas

	var str string = "Texto"
	println(str)

	str2 := "Texto 2"
	println(str2)

	//quando for char utilizar aspas simples e o valor da variável será o correspondente na tabela ASCII

	char := 'B'
	println(char)

	//no Go o erro é um tipo
	var erro error = errors.New("Erro interno")
	println(erro.Error())
}
