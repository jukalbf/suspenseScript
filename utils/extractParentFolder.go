package utils

import "strings"

func ExtractParentFolder(dir string) string {
	parentFolder := strings.Split(dir, "/")
	return parentFolder[len(parentFolder)-2]
}
