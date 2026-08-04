package utils

import (
	"fmt"
	"os"
)

// Cria um arquivo de estilos.
//
// Caso não seja necessário, basta apagar após a criação.
func WriteStylesFile(parentFolder string) {
	fileContent := "// Arquivo de estilos.\n// Se precisar, copie e cole o conteúdo aqui.\n// Se não, apenas delete esse arquivo."
	path := fmt.Sprintf("%s/styles.module.scss", parentFolder)
	err := os.WriteFile(path, []byte(fileContent), 0776)
	CheckErr(err)
}
