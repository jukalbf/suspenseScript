package utils

import (
	"log"
	"os"
	"strings"
)

func ExtractFiles(dirNames []string) ([][]byte, []string) {
	var rawFiles [][]byte
	var filesDir []string

	for _, dir := range dirNames {
		if strings.Contains(dir, ".") {
			data, err := os.ReadFile(dir)

			if err != nil {
				log.Fatal(err)
			}

			rawFiles = append(rawFiles, data)
			filesDir = append(filesDir, dir)
		}
	}

	return rawFiles, filesDir
}
