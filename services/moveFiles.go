package services

import (
	"main/utils"
)

func MoveFiles(path string) (filesDir []string, componentsFolders []string) {
	dirNames := utils.ReadDir(path)
	targetFiles, filesDir := utils.ExtractFiles(dirNames)

	componentsFolders = utils.CreateComponentFolders(filesDir)

	utils.WriteFiles(targetFiles, componentsFolders)

	return
}
