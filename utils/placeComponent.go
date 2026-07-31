package utils

import (
	"fmt"
	"strings"
)

func PlaceComponent(componentFolder string, line *string) {
	extractedFolderName := ExtractParentFolder(componentFolder)
	formatedTagName := fmt.Sprintf("%s%s", strings.ToUpper(extractedFolderName[:1]), extractedFolderName[1:])
	*line = fmt.Sprintf("\t\t\t\t<%s />\n", formatedTagName)
}
