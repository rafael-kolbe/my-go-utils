package input

import (
	"fmt"

	"github.com/chzyer/readline"
)

func Input(msg string) string {
	rl, err := readline.New(msg)
	if err != nil {
		panic(fmt.Sprintf("Erro ao criar readline: %v", err))
	}
	defer rl.Close()

	input, err := rl.Readline()
	if err != nil {
		panic(fmt.Sprintf("Erro ao ler o input: %v", err))
	}

	return input
}