package main

import (
	"github.com/joho/godotenv"              // carrega variáveis do arquivo .env
	"github.com/rs/zerolog"                 // biblioteca de log estruturado
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"      // rotação de arquivos de log
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/devbymarcos/painel-monitoramento/internal/api"
)

//
// ==== MIDDLEWARES ====
//

// Middleware CORS (liberando tudo em DEV, ajuste depois em produção se precisar)
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Middleware de logging (zera cada request recebido)
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		// log estruturado
		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.Path).
			Dur("duration", time.Since(start)).
			Msg("Request processado")
	})
}

//
// ==== MAIN ====
//

func main() {
	// === Carregar variáveis do arquivo .env ===
	_ = godotenv.Load(".env") // se não existir, ignora sem erro

	// === Ler variáveis do .env (ou do sistema) ===
	appMode := os.Getenv("APP_MODE") // debug | production
	port := os.Getenv("PORT")        // porta definida no .env

	if port == "" {
		port = "8080" // valor padrão
	}

	// === Configurar saída de logs ===
	var writer io.Writer

	if appMode == "production" {
		// Em produção → logs só em arquivo, com rotação
		writer = &lumberjack.Logger{
			Filename:   "./Log/app.log", // arquivo base
			MaxSize:    10,              // tamanho máx. em MB antes de rotacionar
			MaxBackups: 5,               // quantos arquivos antigos manter
			MaxAge:     30,              // dias até apagar arquivos antigos
			Compress:   true,            // compacta arquivos antigos (.gz)
		}
	} else {
		// Em debug/dev → logs só no console (mais legível)
		writer = zerolog.ConsoleWriter{Out: os.Stderr}
	}

	// Configura o zerolog
	log.Logger = zerolog.New(writer).With().Timestamp().Logger()

	// === Definir handlers da API ===
	http.HandleFunc("/api/monitor", api.MonitorHandler)
	http.HandleFunc("/api/status", api.StatusHandler)

	// === Servir build do React ===
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("./build", r.URL.Path)

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			// Se não existir o arquivo, serve o index.html (SPA)
			http.ServeFile(w, r, "./build/index.html")
			return
		}

		// Caso contrário, serve o arquivo normalmente (JS, CSS, imagens...)
		http.FileServer(http.Dir("./build")).ServeHTTP(w, r)
	})

	// === Mensagem inicial ===
	log.Info().Msgf("Servidor rodando em http://localhost:%s (modo: %s)", port, appMode)

	// === Encadear middlewares (CORS + Logging) ===
	handler := corsMiddleware(loggingMiddleware(http.DefaultServeMux))

	// === Subir servidor ===
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal().Err(err).Msg("Erro ao iniciar servidor")
	}
}
