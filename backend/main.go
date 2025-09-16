package main

import (
	// variáveis de ambiente
	"net/http"

	"github.com/kardianos/service" // rodar como serviço Windows
	"github.com/rs/zerolog/log"

	"os"

	"github.com/devbymarcos/painel-monitoramento/internal/config"
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
	// execDir não é mais usado para carregar config.json diretamente aqui, mas sim em LoadConfig
	// execDir ainda é necessário para server.SetupServer e caminhos relativos.
	execDir := utils.GetExecDir()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Erro ao carregar configurações")
	}

	port := cfg.Port

	handler := server.SetupServer(execDir, cfg)

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
	// execDir não é mais usado para carregar config.json diretamente aqui, mas sim em LoadConfig
	// execDir ainda é necessário para outras operações, como o serviço.


	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Erro ao carregar configurações")
	}

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
	if cfg.AppMode == "debug" {
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
