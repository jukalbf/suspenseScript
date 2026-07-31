package utils

import "strings"

func ExtractParentFolder(dir string) string {
	splitedFolders := strings.Split(dir, "/")
	return splitedFolders[len(splitedFolders)-1:][0]
}
