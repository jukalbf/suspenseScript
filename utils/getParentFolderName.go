package utils

import (
	"fmt"
	"strings"
)

// Extrai apenas o nome do diretório
func GetParentFolderName(dir string) (fullFolderName string) {
	splitedDir := strings.Split(dir, "/")
	lastIdx := len(splitedDir) - 1
	targetIdx := len(splitedDir) - 2
	folderName := splitedDir[targetIdx:lastIdx][0]

	fullFolderName = fmt.Sprintf("%sContainer", strings.ReplaceAll(folderName, `/[\W]/`, ""))

	return
}
