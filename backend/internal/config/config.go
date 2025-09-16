package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

type Config struct {
	AppMode      string `json:"app_mode"`
	Port         string `json:"port"`
	LogFile      string `json:"log_file"`
	LogMaxSize   int    `json:"log_max_size"`
	LogMaxBackups int    `json:"log_max_backups"`
	LogMaxAge    int    `json:"log_max_age"`
	LogCompress  bool   `json:"log_compress"`
}

// LoadConfig carrega as configurações do arquivo config.json
// Ele tenta carregar do diretório do executável e, se falhar, do diretório de trabalho atual.
func LoadConfig() (*Config, error) {
	var cfg Config

	execDir, err := os.Executable()
	if err != nil {
		return nil, err
	}
	execDir = filepath.Dir(execDir)
	
	// Tenta carregar do diretório do executável
	configPath := filepath.Join(execDir, "config.json")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Warn().Err(err).Msgf("Não foi possível ler o arquivo de configuração em %s, tentando o diretório de trabalho atual.", configPath)
		
		// Se falhar, tenta carregar do diretório de trabalho atual
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		configPath = filepath.Join(cwd, "config.json")
		data, err = ioutil.ReadFile(configPath)
		if err != nil {
			return nil, err
		}
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	// Definir valores padrão caso não estejam no JSON
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return &cfg, nil
}

