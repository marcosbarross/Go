package main

import "fmt"

//Como não foi colocado nenhum tipo na função, ela aceita qualquer tipo de parâmetro, funcionando de forma genérica
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println("-------------")

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(100),
	}

	fmt.Println(mapa)
}
