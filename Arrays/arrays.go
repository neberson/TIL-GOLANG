package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	nome, idade := devolveNomeEIdade()

	fmt.Println("Olá", nome)
	fmt.Println("Sua idade é", idade)

	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "Neberson Andrade"
	idade := 29
	return nome, idade
}

func exibeIntroducao() {
	var nome string = "Neberson Andrade"
	var versao float32 = 1.1
	var idade int = 29

	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O endereço da váriavel comando é", &comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string

	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://www.google.com"
	sites[2] = "https://www.caelum.com.br"
	sites[3] = "https://www.facebook.com"

	fmt.Println(sites)

	site := "https://www.alura.com.br"

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
