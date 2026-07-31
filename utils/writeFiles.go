package utils

import (
	"fmt"
	"os"
)

func WriteFiles(targetFiles [][]byte, parentsFolder []string) {
	fmt.Println("=> Escrevendo os arquivos dentro das novas pastas...")
	for i, file := range targetFiles {
		fileName := fmt.Sprintf("%s/index.tsx", parentsFolder[i])
		err := os.WriteFile(fileName, file, os.ModePerm)
		WriteStylesFile(parentsFolder[i])
		CheckErr(err)
	}
}
