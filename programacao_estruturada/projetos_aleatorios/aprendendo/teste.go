package main

import (
	"fmt"
	"os"
)

var categoria = []string{}

func main() {
	for {
		menu()
		var escolha int
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			titulo("Cadastro de Categorias")
			adCategoria()
		case 2:
			titulo("Cadastro de Itens")
			if len(categoria) > 0 {
				adItem()
			} else {
				fmt.Println("Ops, sua categoria de compras esta vazia! Cadastre uma.")
			}
		case 3:
			titulo("Remover um item")
		case 4:
			titulo("Remover uma categoria")
			moCateg()

			var numCategoria int
			fmt.Println("Número Categoria: ")
			numCategoria = numCategoria - 1
			fmt.Scan(&numCategoria)

			cateRemover(numCategoria)

		case 5:
			moCateg()
		case 6:

		case 0:
			fmt.Println("Sair")
			os.Exit(-1)
		}
	}
}

func titulo(texto string) {
	fmt.Println("\n===============================")
	fmt.Printf("    %s", texto)
	fmt.Println("\n===============================")
}

func menu() {
	fmt.Println("\n=====================================")
	fmt.Println("            MENU            ")
	fmt.Println("1 - Adicionar a categoria")
	fmt.Println("2 - Adicionar item")
	fmt.Println("3 - Remover um item da lista")
	fmt.Println("4 - Remover um item da categoria")
	fmt.Println("5 - Ver categorias")
	fmt.Println("6 - Ver categorias e seus itens")
	fmt.Println("0 - Sair")
	fmt.Println("=====================================")
	fmt.Print("\nEscolha: ")
}

func adCategoria() {
	fmt.Print("\nNome categoria: ")
	var novaCategoria string
	fmt.Scan(&novaCategoria)
	categoria = append(categoria, novaCategoria)

	fmt.Println("Categoria", novaCategoria, "adicionada com sucesso!")
}

func adItem() {
	fmt.Print("Nome item:")
	var item string
	fmt.Scan(&item)

	tCategorias := len(categoria)

	fmt.Println("Escolha a categoria relacionada: escolha de 1 a", tCategorias)
	moCateg()
	var nCategoria int
	fmt.Scan(&nCategoria)
	nCategoria -= -1
}

func itemRemover(itemRemover int) {

}

func cateRemover(cateRemover int) {
	cateRemover = cateRemover - 1

	antigaCategoria := categoria[cateRemover]

	categoria = append(categoria[:cateRemover], categoria[cateRemover+1:]...)

	fmt.Printf("Categoria %s removida com sucesso!", antigaCategoria)
}

func moCateg() {
	fmt.Println("\nCategorias cadastradas")
	for i := 0; i < len(categoria); i++ {
		fmt.Printf("%d - %s\n", i+1, categoria[i])
	}
}
