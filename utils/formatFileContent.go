package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Formata conteudo do arquivo, removendo imports e consts.
func FormatFileContent(scanner *bufio.Scanner, targetFile *os.File) (formatedRows []string) {
	for scanner.Scan() {
		text := scanner.Text()

		if !isImport(text) && !isConst(text) {
			textFormated := fmt.Sprintf("%s\n", text)
			formatedRows = append(formatedRows, textFormated)
		}
	}

	return
}

// Verifica se a linha(text) atual contém import de algo client-side.
func isImport(text string) bool {
	containsImport := strings.Contains(text, "import")
	containsUse := strings.Contains(text, "use")
	containsParam := strings.Contains(text, "params")
	hasUse := containsImport && containsUse
	hasParam := containsImport && containsParam

	return hasUse || hasParam
}

// Verifica se a linha(text) atual contém uma const client-side.
func isConst(text string) bool {
	hasConst := strings.Contains(text, "const")
	isArrowFunc := strings.Contains(text, "() =>")
	hasExport := strings.Contains(text, "export")

	return hasConst && !hasExport && !isArrowFunc
}
