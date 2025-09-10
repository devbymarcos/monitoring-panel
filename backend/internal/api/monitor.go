package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func MonitorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	dataDir := filepath.Join(".", "Data")
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(dataDir, os.ModePerm)
	}

	filePath := filepath.Join(dataDir, "data.json")
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Erro ao processar JSON", http.StatusInternalServerError)
		return
	}
	err = os.WriteFile(filePath, jsonBytes, 0644)
	if err != nil {
		http.Error(w, "Erro ao salvar arquivo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"sucesso"}`))
}
