package main

import (
	"net/http"

	"github.com/joho/godotenv"     // variáveis de ambiente
	"github.com/kardianos/service" // rodar como serviço Windows
	"github.com/rs/zerolog/log"

	"os"
	"path/filepath"

	"github.com/devbymarcos/painel-monitoramento/internal/server"
	"github.com/devbymarcos/painel-monitoramento/internal/utils"
)

//
// ==== SERVICE STRUCT ====
//

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	// Carregar variáveis de ambiente sempre do diretório do executável
	execDir := utils.GetExecDir()
	_ = godotenv.Load(filepath.Join(execDir, ".env"))

	appMode := os.Getenv("APP_MODE") // debug | production
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := server.SetupServer(execDir, appMode, port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal().Err(err).Msg("Erro ao iniciar servidor")
	}
}

func (p *program) Stop(s service.Service) error {
	log.Info().Msg("Serviço sendo parado...")
	return nil
}

//
// ==== MAIN ====
//

func main() {
	// Carregar APP_MODE logo no início
	execDir := utils.GetExecDir()
	_ = godotenv.Load(filepath.Join(execDir, ".env"))
	appMode := os.Getenv("APP_MODE")

	svcConfig := &service.Config{
		Name:        "MonitorWork",  // nome interno (sc query PainelDbm)
		DisplayName: "MonitorWork", // nome visível no services.msc
		Description: "Serviço MonitorWork que roda o painel de monitoramento com API e React SPA",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Erro ao criar serviço")
	}

	// Se APP_MODE=debug → roda direto no console
	if appMode == "debug" {
		log.Info().Msg("Rodando em modo console (debug)")
		prg.run()
		return
	}

	// Se tiver argumentos (install/start/stop/uninstall)
	if len(os.Args) > 1 {
		if err := service.Control(s, os.Args[1]); err != nil {
			log.Fatal().Err(err).Msg("Erro ao controlar serviço")
		}
		return
	}

	// APP_MODE=production → roda como serviço
	if err := s.Run(); err != nil {
		log.Fatal().Err(err).Msg("Erro ao executar serviço")
	}
}
