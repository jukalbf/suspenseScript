package modules

import (
	"bufio"
	"fmt"
	"main/utils"
	"os"
	"strings"
)

func RewriteFiles(filesDir []string, componentFolders []string) {
	file, err := os.Open(filesDir[1])

	utils.CheckErr(err)

	scanner := bufio.NewScanner(file)

	lines := utils.ExcludeClientImports(scanner)

	for i, line := range lines {
		if i == 0 {
			extractedFolderName := utils.ExtractParentFolder(componentFolders[1])
			componentName := fmt.Sprintf("%s%s", strings.ToUpper(extractedFolderName[:1]), extractedFolderName[1:])
			lines[i] = fmt.Sprintf("import { %s } from '@/components/%s'\n", componentName, extractedFolderName)
		}

		if strings.Contains(line, "<PageHeader") {
			utils.PlaceComponent(componentFolders[1], &lines[i+1])
		}
	}

	utils.WriteSingleFile(filesDir[1], strings.Join(lines, ""))
}
