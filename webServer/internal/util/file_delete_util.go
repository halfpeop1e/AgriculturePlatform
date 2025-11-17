package util

import (
	"go-agriculture/internal/config"
	"os"
)

func FileDeleteUtil(fileName string, fileType FileType) {
	if fileName == "" || fileName == "default.png" {
		return
	}
	var beforePath string
	if fileType == IsFileType {
		beforePath = config.GetConfig().StaticFilePath
	} else {
		beforePath = config.GetConfig().StaticAvatarPath
	}
	filePath := beforePath + "/" + fileName
	os.Remove(filePath)
}
