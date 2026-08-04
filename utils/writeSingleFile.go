package utils

import (
	"os"
)

// Cria novo arquivo.
func WriteSingleFile(dir string, content string) {
	os.WriteFile(dir, []byte(content), os.ModePerm)
}
