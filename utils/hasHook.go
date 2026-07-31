package utils

import (
	"os"
	"strings"
)

// Verifica os arquivos que contem a função `useSearchParams()`.
func HasHook(rawFile []byte) bool {
	fileContent := string(rawFile)
	expression := os.Args[1]

	return strings.Contains(fileContent, expression)
}
