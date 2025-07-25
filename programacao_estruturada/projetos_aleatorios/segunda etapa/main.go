package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Estrutura para receber a resposta da API
type APIResponse struct {
	Results []Character `json:"results"`
}

// Estrutura para guardar os dados de cada personagem
type Character struct {
	Name   string `json:"name"`
	Image  string `json:"image"`
	Status string `json:"status"`
}

// Nosso template HTML para exibir os personagens
var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", characterHandler)

	fmt.Println("Portal Rick and Morty rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func characterHandler(w http.ResponseWriter, r *http.Request) {
	// URL da API do Rick and Morty
	url := "https://rickandmortyapi.com/api/character"

	// Fazendo a "ligação" (request) para a API
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Wubba Lubba Dub Dub! O portal falhou.", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decodificando a resposta (que vem em JSON)
	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		http.Error(w, "Falha ao traduzir os dados da outra dimensão.", http.StatusInternalServerError)
		return
	}

	// Enviando os dados para o nosso HTML
	tmpl.Execute(w, apiResponse)
}
