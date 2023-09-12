package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo") // Para usar 1 - goroutines, basta colocar a palavra go antes da chamada da função
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
