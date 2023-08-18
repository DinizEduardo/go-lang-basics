package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Eduardo"
	idade := 23
	versao := 1.1
	fmt.Println("Olá, sr.", nome, " sua idade é", idade)
	fmt.Println("Esse programa está na versão", versao)

	fmt.Println("O tipo da variavel nome é: ", reflect.TypeOf((versao)))
}
