package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Caminho do arquivo dentro da pasta Data
	dataPath := filepath.Join(".", "Data", "data.json")

	// Lê o arquivo
	file, err := os.ReadFile(dataPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler Data/data.json: %v", err), http.StatusInternalServerError)
		return
	}

	// Define cabeçalho e escreve resposta
	w.Header().Set("Content-Type", "application/json")
	w.Write(file)
}
