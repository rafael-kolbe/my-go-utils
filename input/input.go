// Package input fornece funções para ler entrada do usuário com uma mensagem de prompt.
package input

import (
	"log"

	"github.com/chzyer/readline"
)

// Lê a entrada do usuário exibindo uma mensagem de prompt.
// A função retorna a string digitada pelo usuário.
func Prompt(msg string) string {
	rl, err := readline.New(msg)
	if err != nil {
		log.Fatal("Erro ao criar readline: ", err)
	}
	defer rl.Close()

	prompt, err := rl.Readline()
	if err != nil {
		log.Fatal("Erro ao ler o prompt: ", err)
	}

	return prompt
}