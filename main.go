package main

import (
	"fmt"
	"log"
	"main/utils"
	"os"
)

func main() {
	fmt.Println("=> Iniciando processo...")
	// codePath := `/home/jdev/Documents/projects/gcp_frontend/src/app/(pages)/`
	codePath := `/home/jdev/Documents/projects/gcp_frontend/src/app/(pages)/`
	var parentsFolder []string

	dirNames := utils.ReadDir(codePath)
	rawFiles, filesDir := utils.ExtractFiles(dirNames)

	count, targetFiles := utils.CountCheck(rawFiles)

	for _, dir := range filesDir {
		parent := utils.ExtractParentFolder(dir)
		targetDir := fmt.Sprintf("src/components/%sContainer", parent)
		os.Mkdir(targetDir, 0777)
		parentsFolder = append(parentsFolder, targetDir)
	}

	for i, file := range targetFiles {
		fileName := fmt.Sprintf("%s/index.tsx", parentsFolder[i])
		err := os.WriteFile(fileName, file, 0777)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(count)
}

// Log dos diretorios. `defaultPath` — Caminho padrão
func showDirs(dirNames []string, defaultPath string) {
	for _, dir := range dirNames {
		fmt.Printf("| ==> %s%s\n", defaultPath, dir)
	}
}
