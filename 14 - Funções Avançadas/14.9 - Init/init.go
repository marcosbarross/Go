package main

import "fmt"

//função init é executada antes da função main
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
