package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Retorno do json com nome e horário
type ProjetoKorpResponse struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func main() {
	// endponint GET /projeto-korp
	http.HandleFunc("/projeto-korp", RouteProjeto)
	http.ListenAndServe(":8080", nil) // escutando a porta 8080
}

func RouteProjeto(w http.ResponseWriter, r *http.Request) {
	// Garante que apenas o endpoint GET /projeto-korp seja aceito
	if r.URL.Path != "/projeto-korp" || r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Define o nome e o horario com o formato UTC
	response := ProjetoKorpResponse{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}