package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/devbymarcos/painel-monitoramento/internal/api"
)

func main() {
	// Endpoint da API
	http.HandleFunc("/api/monitor", api.MonitorHandler)
	http.HandleFunc("/api/status", api.StatusHandler)

	// Servir arquivos do React (SPA)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Caminho do arquivo solicitado
		path := filepath.Join("./build", r.URL.Path)

		// Verificar se o arquivo existe
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			// Se não existir, devolve o index.html (SPA)
			http.ServeFile(w, r, "./build/index.html")
			return
		}

		// Se existir, serve normalmente
		http.FileServer(http.Dir("./build")).ServeHTTP(w, r)
	})

	// Porta
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando em http://localhost:%s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
