package main

import (
	"fmt"
	"main/modules"
	"main/utils"
)

func main() {
	utils.CheckArgs()

	fmt.Println("=> Iniciando processo...")

	codePath := `/home/jdev/Documents/projects/gcp_frontend/src/app/(pages)/`

	filesDir, componentsFolders := modules.MoveFiles(codePath)

	modules.RewriteFiles(filesDir, componentsFolders)
}

// Log dos diretorios. `defaultPath` — Caminho padrão
func showDirs(dirNames []string, defaultPath string) {
	for _, dir := range dirNames {
		fmt.Printf("| ==> %s%s\n", defaultPath, dir)
	}
}
