package main

import (
	"fmt"
	"main/modules"
	"main/utils"
)

func main() {
	utils.CheckArgs()

	fmt.Println("=> Iniciando processo...")

	appPath := `/home/jdev/Documents/suspenseScript/examples/src/app/(pages)/`

	filesDir, componentsFolders := modules.MoveFiles(appPath)

	modules.RewriteFiles(filesDir, componentsFolders)
}
