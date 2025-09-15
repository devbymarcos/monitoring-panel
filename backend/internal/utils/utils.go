package utils

import (
	"os"
	"path/filepath"
)

// Retorna o diretório onde o executável está rodando
func GetExecDir() string {
	execPath, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(execPath)
}
