package utils

import (
	"fmt"
	"strings"
)

// Implementa o novo componente no arquivo original.
func PlaceComponent(componentFolder string, targetRow *string) {
	extractedFolderName := ExtractParentFolder(componentFolder)
	formatedTagName := fmt.Sprintf("%s%s", strings.ToUpper(extractedFolderName[:1]), extractedFolderName[1:])
	*targetRow = fmt.Sprintf("\t\t\t\t<%s />\n", formatedTagName)
}
