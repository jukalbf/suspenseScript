package utils

import (
	"log"
	"os"
)

// Lê diretório(dir) base e retorna um slice com os diretórios filhos — apenas diretórios
func ReadDir(dir string) []string {
	files, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	dirNames := extractDirs(files, dir)

	return dirNames
}

// Extrai os diretórios para um slice
func extractDirs(files []os.DirEntry, dir string) []string {
	var dirNames []string

	for _, file := range files {
		if file.Name() != "" {
			fullDir := dir + file.Name()
			dirNames = append(dirNames, fullDir)

			if file.IsDir() {
				subDir := ReadDir(dir + file.Name() + "/")
				dirNames = append(dirNames, subDir...)
			}
		}
	}

	return dirNames
}
