package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // WaitGroup é uma estrutura que permite que a goroutine principal espere o término de todas as 1 - goroutines criadas
	waitGroup.Add(2)             // Adiciona a quantidade de 1 - goroutines que serão criadas

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() // Ao finalizar a goroutine, é necessário chamar o método Done() para informar que a goroutine foi finalizada
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // Aguarda o término de todas as 1 - goroutines
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
