package utils

import (
	"os"
)

func WriteSingleFile(dir string, content string) {
	os.WriteFile(dir, []byte(content), os.ModePerm)
}
