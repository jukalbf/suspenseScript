package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func ExcludeClientImports(scanner *bufio.Scanner) (lines []string) {
	for scanner.Scan() {
		text := scanner.Text()
		containsImport := strings.Contains(text, "import")
		containsUse := strings.Contains(text, "use")
		containsProvider := strings.Contains(text, "Provider")
		containsParam := strings.Contains(text, "params")
		containsCapParam := strings.Contains(text, "Params")
		containsConst := strings.Contains(text, "const") && !strings.Contains(text, "export")
		hasUse := containsImport && containsUse
		hasProvider := containsImport && containsProvider

		if !hasUse && !hasProvider && !containsUse && !containsParam && !containsProvider && !containsCapParam && !containsConst {
			stringFormat := fmt.Sprintf("%s\n", text)
			lines = append(lines, stringFormat)
		}
	}

	return
}
