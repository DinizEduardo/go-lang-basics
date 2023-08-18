package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	exibeNomes()
	exibeIntroducao()
	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Eduardo"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://caelum.com.br"
	sites[2] = "https://google.com.br"
	sites[3] = "https://youtube.com.br"
	fmt.Println(reflect.TypeOf(sites))
	site := sites[0]

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "está com problemas! Status code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Eduardo", "Daniel", "Bernardo"}
	nomes = append(nomes, "Aparecida")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem ", len(nomes))
}
