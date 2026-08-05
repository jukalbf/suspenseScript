package modules

import (
	"bufio"
	"fmt"
	"main/utils"
	"os"
	"strings"
)

// Módulo responsável por reescrever arquivos principais.
func RewriteFiles(filesDir []string, componentFolders []string) {
	for _, file := range filesDir {
		targetFile, err := os.Open(file)

		utils.CheckErr(err)

		scanner := bufio.NewScanner(targetFile)

		fmt.Println("=> Removendo imports e consts...")
		formatedRows := utils.FormatFileContent(scanner, targetFile)

		fmt.Println("=> Imports e consts removidos com sucesso.")

		for _, component := range componentFolders {
			fmt.Println("------------------------------------------")
			fmt.Printf("===> Componente atual: %s\n", component)

			for j, row := range formatedRows {
				if j == 0 {
					insertComponentImport(component, &formatedRows[j])
				}

				if rowHasPageHeader(row) {
					fmt.Printf("=====> Inserindo componente em: %s\n", file)
					utils.PlaceComponent(component, &formatedRows[j+1])
				}

				utils.WriteSingleFile(file, strings.Join(formatedRows, ""))
			}
		}
	}
}

// Insere o import do novo componente.
func insertComponentImport(component string, row *string) {
	extractedFolderName := utils.ExtractParentFolder(component)
	componentName := fmt.Sprintf("%s%s", strings.ToUpper(extractedFolderName[:1]), extractedFolderName[1:])
	*row = fmt.Sprintf("import %s from '@/components/%s'\n", componentName, extractedFolderName)
}

// Verifica se a linha(row) atual contem a tag `<PageHeader />`
func rowHasPageHeader(row string) bool {
	return strings.Contains(row, "<PageHeader")
}
