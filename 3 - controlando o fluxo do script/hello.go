package main

import (
	"fmt"
)

func main() {
	nome := "Eduardo"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo...")
	} else {
		fmt.Println("Não conheço este comando")
	}
}
