package utils

import (
	"strings"
)

// Extrai o nome da pasta e transforma a primeira letra em maiúscula.
func ExtractParentFolder(dir string) string {
	splitedFolders := strings.Split(dir, "/")
	return splitedFolders[len(splitedFolders)-1:][0]
}
