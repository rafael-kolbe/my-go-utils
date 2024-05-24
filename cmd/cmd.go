// Package cmd fornece funções para executar comandos no sistema operacional.
package cmd

import (
	"os"
	"os/exec"
	"strings"
)

// Executa um comando no sistema operacional.
// O comando é especificado como uma string, e o diretório de trabalho pode ser especificado.
// Se o diretório for uma string vazia, o comando será executado no diretório atual.
func RunCommand(command string, dir string) error {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if dir != "" {
		cmd.Dir = dir
	}
	return cmd.Run()
}