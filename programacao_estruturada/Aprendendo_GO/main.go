package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {
	apresentacao()

	for {
		exibeMenu()

		var escolha int
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Ops, opção inválida!")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Ixibir Logs")
	fmt.Println("0 - Sair")
	fmt.Println(" ")
}

func apresentacao() {
	nome := "Douglas"
	fmt.Println("Olá, sr.", nome)
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := lerSitesArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site:", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println(" ")
	}

	fmt.Println(" ")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro!", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site: ", site, "esta com problemas!", resp.StatusCode)
	}
}

func lerSitesArquivo() []string {

	var sites []string

	//arquivo, err := os.Open("sites.txt")
	arquivo, err := os.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro!", err)
	}

	fmt.Println(string(arquivo))
	return sites
}
