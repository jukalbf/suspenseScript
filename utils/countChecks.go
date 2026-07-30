package utils

import (
	"strings"
)

func CountCheck(rawFile [][]byte) (int, [][]byte) {
	count := 0
	var targetFiles [][]byte

	for _, file := range rawFile {
		rawTextFileContent := string(file)

		if strings.Contains(rawTextFileContent, "useSearchParams()") {
			count++
			targetFiles = append(targetFiles, file)
		}
	}

	return count, targetFiles
}
