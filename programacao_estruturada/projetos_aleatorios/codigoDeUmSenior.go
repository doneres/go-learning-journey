package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// --- 1. Estruturas de Dados Claras (Structs) ---
// Usar structs torna o modelo de dados explícito e extensível.
// Se um item precisar de um preço ou quantidade no futuro, é fácil adicionar.
type Item struct {
	Nome string
}

type Categoria struct {
	Nome  string
	Itens []Item
}

// --- Main: O Controlador Principal ---
// A função main orquestra o programa. Ela detém o estado (a lista)
// e chama outras funções para modificá-lo ou exibi-lo.
func main() {
	// O estado da aplicação vive aqui, não em variáveis globais.
	listaDeCompras := []Categoria{}

	// Dados iniciais para facilitar o teste
	listaDeCompras = append(listaDeCompras, Categoria{Nome: "Bebidas", Itens: []Item{{Nome: "Água"}, {Nome: "Suco de Laranja"}}})
	listaDeCompras = append(listaDeCompras, Categoria{Nome: "Padaria", Itens: []Item{{Nome: "Pão Francês"}}})

	for {
		menu()
		escolha := lerInteiro("Sua escolha: ")

		switch escolha {
		case 1:
			titulo("Adicionar Nova Categoria")
			// Passamos um ponteiro para que a função possa modificar o slice original.
			adicionarCategoria(&listaDeCompras)
		case 2:
			titulo("Adicionar Novo Item")
			adicionarItem(listaDeCompras)
		case 3:
			titulo("Remover Categoria")
			// A função de remoção retorna o novo slice, que reatribuímos.
			listaDeCompras = removerCategoria(listaDeCompras)
		case 4:
			titulo("Remover Item")
			removerItem(listaDeCompras)
		case 5:
			titulo("Minha Lista de Compras")
			mostrarTudo(listaDeCompras)
		case 0:
			fmt.Println("\nAté mais!")
			os.Exit(0)
		default:
			fmt.Println("\nErro: Opção inválida. Tente novamente.")
		}
		pressioneEnterParaContinuar()
	}
}

// --- Funções de Lógica de Negócio ---

func adicionarCategoria(lista *[]Categoria) {
	nome := lerString("Digite o nome da nova categoria: ")
	if nome == "" {
		fmt.Println("Erro: O nome não pode ser vazio.")
		return
	}

	// Verifica se a categoria já existe (ignorando maiúsculas/minúsculas)
	for _, c := range *lista {
		if strings.EqualFold(c.Nome, nome) {
			fmt.Printf("Erro: A categoria '%s' já existe.\n", nome)
			return
		}
	}

	novaCategoria := Categoria{Nome: nome, Itens: []Item{}}
	*lista = append(*lista, novaCategoria)
	fmt.Printf("Categoria '%s' adicionada com sucesso.\n", nome)
}

func adicionarItem(lista []Categoria) {
	if len(lista) == 0 {
		fmt.Println("Nenhuma categoria cadastrada. Adicione uma categoria primeiro.")
		return
	}

	fmt.Println("Em qual categoria deseja adicionar o item?")
	indiceCat, sucesso := selecionarCategoria(lista)
	if !sucesso {
		return // Usuário cancelou ou a escolha foi inválida
	}

	nomeItem := lerString("Digite o nome do novo item: ")
	if nomeItem == "" {
		fmt.Println("Erro: O nome do item não pode ser vazio.")
		return
	}

	// Adiciona o novo item ao slice de Itens da categoria escolhida.
	// Note que estamos modificando um elemento dentro do slice `lista`,
	// por isso não precisamos de um ponteiro para o slice inteiro aqui.
	lista[indiceCat].Itens = append(lista[indiceCat].Itens, Item{Nome: nomeItem})
	fmt.Printf("Item '%s' adicionado à categoria '%s'.\n", nomeItem, lista[indiceCat].Nome)
}

func removerCategoria(lista []Categoria) []Categoria {
	if len(lista) == 0 {
		fmt.Println("Nenhuma categoria para remover.")
		return lista
	}

	fmt.Println("Qual categoria você deseja remover?")
	indice, sucesso := selecionarCategoria(lista)
	if !sucesso {
		return lista // Retorna a lista original se a seleção falhar
	}

	categoriaRemovida := lista[indice].Nome
	// O idioma padrão para remover um item de um slice.
	novaLista := append(lista[:indice], lista[indice+1:]...)
	fmt.Printf("Categoria '%s' e todos os seus itens foram removidos.\n", categoriaRemovida)
	return novaLista
}

func removerItem(lista []Categoria) {
	if len(lista) == 0 {
		fmt.Println("Nenhuma categoria cadastrada.")
		return
	}

	fmt.Println("De qual categoria você deseja remover um item?")
	indiceCat, sucesso := selecionarCategoria(lista)
	if !sucesso {
		return
	}

	categoriaSelecionada := &lista[indiceCat] // Pega um ponteiro para a categoria para facilitar a modificação
	if len(categoriaSelecionada.Itens) == 0 {
		fmt.Printf("A categoria '%s' não possui itens para remover.\n", categoriaSelecionada.Nome)
		return
	}

	fmt.Println("Qual item você deseja remover?")
	for i, item := range categoriaSelecionada.Itens {
		fmt.Printf("%d - %s\n", i+1, item.Nome)
	}

	escolha := lerInteiro("Sua escolha: ")
	indiceItem := escolha - 1

	if indiceItem < 0 || indiceItem >= len(categoriaSelecionada.Itens) {
		fmt.Println("Erro: Número do item inválido.")
		return
	}

	itemRemovido := categoriaSelecionada.Itens[indiceItem].Nome
	categoriaSelecionada.Itens = append(categoriaSelecionada.Itens[:indiceItem], categoriaSelecionada.Itens[indiceItem+1:]...)
	fmt.Printf("Item '%s' removido da categoria '%s'.\n", itemRemovido, categoriaSelecionada.Nome)
}

// --- Funções de UI e Helpers ---

func menu() {
	fmt.Println("\n=====================================")
	fmt.Println("      MENU - LISTA DE COMPRAS      ")
	fmt.Println("=====================================")
	fmt.Println("1 - Adicionar Categoria")
	fmt.Println("2 - Adicionar Item")
	fmt.Println("3 - Remover Categoria")
	fmt.Println("4 - Remover Item")
	fmt.Println("5 - Ver Lista Completa")
	fmt.Println("0 - Sair")
	fmt.Println("-------------------------------------")
}

func titulo(t string) {
	fmt.Printf("\n---[ %s ]---\n\n", strings.ToUpper(t))
}

func mostrarTudo(lista []Categoria) {
	if len(lista) == 0 {
		fmt.Println("Sua lista de compras está vazia.")
		return
	}

	// Ordena as categorias para exibição consistente
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].Nome < lista[j].Nome
	})

	for _, categoria := range lista {
		fmt.Printf("## %s\n", categoria.Nome)
		if len(categoria.Itens) == 0 {
			fmt.Println("   (vazio)")
		} else {
			for _, item := range categoria.Itens {
				fmt.Printf("   - %s\n", item.Nome)
			}
		}
	}
}

// selecionarCategoria é uma função helper que encapsula a lógica de
// mostrar a lista de categorias e capturar uma escolha válida do usuário.
// Retorna o índice (base-0) e um booleano indicando sucesso.
func selecionarCategoria(lista []Categoria) (int, bool) {
	for i, c := range lista {
		fmt.Printf("%d - %s\n", i+1, c.Nome)
	}
	escolha := lerInteiro("Sua escolha (ou 0 para cancelar): ")

	if escolha == 0 {
		fmt.Println("Operação cancelada.")
		return 0, false
	}

	indice := escolha - 1

	if indice < 0 || indice >= len(lista) {
		fmt.Println("Erro: Número de categoria inválido.")
		return 0, false
	}
	return indice, true
}

// --- Funções de Input Robustas ---

// lerString lê uma linha inteira do console, ao contrário do Scan que para no espaço.
func lerString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

// lerInteiro garante que o usuário digite um número válido, repetindo o prompt se necessário.
func lerInteiro(prompt string) int {
	for {
		texto := lerString(prompt)
		numero, err := strconv.Atoi(texto)
		if err == nil {
			return numero
		}
		fmt.Println("Erro: Por favor, digite um número válido.")
	}
}

func pressioneEnterParaContinuar() {
	lerString("\nPressione ENTER para voltar ao menu...")
}
