package utils

import (
	"log"
	"os"
)

// Verifica se o argumento foi passado.
func CheckArgs() {
	expression := os.Args

	if len(expression) != 2 {
		log.Fatal("Script deve ser rodado com pelo menos e no máximo um argumento.")
	}
}
