package utils

import (
	"fmt"
	"os"
	"strings"
)

// Extrai os arquivos que tem o hook target.
func ExtractFiles(dirNames []string) (rawFiles [][]byte, filesDir []string) {
	fmt.Println("=> Extraindo arquivo que contem a expressão")

	for _, dir := range dirNames {
		if strings.Contains(dir, ".") {
			data, err := os.ReadFile(dir)

			CheckErr(err)

			if HasHook(data) {
				rawFiles = append(rawFiles, data)
				filesDir = append(filesDir, dir)
			}
		}
	}

	return
}
