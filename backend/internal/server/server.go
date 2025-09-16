package server

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/devbymarcos/painel-monitoramento/internal/api"
	"github.com/devbymarcos/painel-monitoramento/internal/config"
	"github.com/devbymarcos/painel-monitoramento/internal/middleware"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupServer(execDir string, cfg *config.Config) http.Handler {
	// Configuração de logs
	var writer io.Writer
	if cfg.AppMode == "production" {
		logDir := filepath.Join(execDir, "Log")
		_ = os.MkdirAll(logDir, os.ModePerm)

		// Logs em JSON com rotação
		writer = &lumberjack.Logger{
			Filename:   filepath.Join(execDir, cfg.LogFile),
			MaxSize:    cfg.LogMaxSize, // MB
			MaxBackups: cfg.LogMaxBackups,
			MaxAge:     cfg.LogMaxAge,   // dias
			Compress:   cfg.LogCompress, // compacta logs antigos
		}
	} else {
		// Logs legíveis no console (debug)
		writer = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"}
	}
	log.Logger = zerolog.New(writer).With().Timestamp().Logger()

	// Handlers da API
	http.HandleFunc("/api/monitor", api.MonitorHandler)
	http.HandleFunc("/api/status", api.StatusHandler)

	// Servir arquivos do React (SPA)
	buildDir := filepath.Join(execDir, "build")
	fs := http.FileServer(http.Dir(buildDir))
	http.Handle("/static/", fs) // assets (js, css, imagens)

	// SPA fallback → sempre devolver index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(buildDir, r.URL.Path)
		info, err := os.Stat(path)
		if err == nil && !info.IsDir() {
			http.ServeFile(w, r, path)
			return
		}
		http.ServeFile(w, r, filepath.Join(buildDir, "index.html"))
	})

	// Encadear middlewares
	handler := middleware.CorsMiddleware(middleware.LoggingMiddleware(http.DefaultServeMux))

	log.Info().Msgf("Servidor rodando em http://localhost:%s (modo: %s)", cfg.Port, cfg.AppMode)
	// The server will be started in the run function.
	return handler
}
