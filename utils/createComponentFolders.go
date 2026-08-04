package utils

import (
	"fmt"
	"os"
)

// Cria as pastas dos componentes.
func CreateComponentFolders(filesDir []string) (componentsDir []string) {
	fmt.Println("=> Criando pastas para o componentes...")

	for _, dir := range filesDir {
		fullFolderName := GetParentFolderName(dir)
		parentDir := fmt.Sprintf("/home/user/path/src/components/%s", fullFolderName)
		os.Mkdir(parentDir, os.ModePerm)
		componentsDir = append(componentsDir, parentDir)
	}

	fmt.Println("=> Pastas criadas.")

	return componentsDir
}
