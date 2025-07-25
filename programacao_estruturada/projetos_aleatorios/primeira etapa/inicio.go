package main

import (
	"fmt"
)

func main() {

	fmt.Println("===================================")
	fmt.Println("      A CAVERNA PERDIDA")
	fmt.Println("===================================")
	fmt.Println("\nVocê acorda em um lugar úmido e escuro. A cabeça dói um pouco.")
	fmt.Println("Ao se levantar, você percebe que está na entrada de uma caverna.")
	fmt.Println("À sua frente, a caverna se divide em dois túneis.")
	fmt.Println("\nO que você faz?")
	fmt.Println("1. Seguir pelo túnel da esquerda, de onde vem um brilho fraco.")
	fmt.Println("2. Seguir pelo túnel da direita, que é escuro e silencioso.")

	var escolha int
	fmt.Print("\nDigite sua escolha (1 ou 2): ")
	fmt.Scan(&escolha)

	fmt.Scan(&escolha)

	if escolha == 1 {
		fmt.Println("\nVocê segue o brilho e encontra uma câmara cheia de cristais reluzentes.")
		fmt.Println("No centro, há um baú. Você o abre e encontra um mapa antigo!")
		fmt.Println("FIM - Você encontrou um tesouro!")

	} else if escolha == 2 {
		fmt.Println("\nVocê entra no túnel escuro. Após alguns passos, ouve um barulho alto.")
		fmt.Println("Um morcego gigante voa em sua direção! Você corre para fora da caverna, assustado.")
		fmt.Println("FIM - Você escapou por pouco!")

	} else {
		fmt.Println("\nOpção inválida! O medo te paralisa e você não entra em nenhum túnel.")
		fmt.Println("FIM - Você ficou parado para sempre.")
	}

	fmt.Println("\n===================================")
}
