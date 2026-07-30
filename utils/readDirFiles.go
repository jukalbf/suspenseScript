package utils

import (
	"log"
	"os"
)

// Lê todo o conteúdo do diretório
func ReadDirFiles(dir string) []string {
	files, err := os.ReadDir(dir)
	var filesNames []string

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filesNames = append(filesNames, file.Name())
	}

	return filesNames
}
