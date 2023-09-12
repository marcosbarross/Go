package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //chan vem de channel
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal //recebendo dados do canal, a segunda varíavel é um booleano que indica se o canal está aberto ou não gerado automaticamente ao ser declarada
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	//Outro modo de ler dados de um canal
	for mensagem := range canal { //range é usado para ler dados de um canal
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //enviando dados para o canal
		time.Sleep(time.Second)
	}

	close(canal)
}
