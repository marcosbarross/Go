package main

import "fmt"

func main() {
	canal := make(chan string, 2) //o segundo parametro é a capacidade do buffer

	canal <- "Olá mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
