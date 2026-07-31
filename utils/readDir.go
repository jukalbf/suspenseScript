package utils

import (
	"os"
)

// Lê diretório(dir) base e retorna um slice com os diretórios filhos — apenas diretórios
func ReadDir(dir string) (dirNames []string) {
	files, err := os.ReadDir(dir)

	CheckErr(err)

	dirNames = extractDirs(files, dir)

	return
}

// Extrai os diretórios para um slice
func extractDirs(files []os.DirEntry, dir string) (dirNames []string) {
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

	return
}
