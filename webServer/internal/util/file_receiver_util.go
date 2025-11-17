package util

import (
	"go-agriculture/internal/config"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileType int

const (
	IsFileType FileType = 0

	IsAvatarType FileType = 1
)

func FileReceiverUtil(files []*multipart.FileHeader, fileType FileType) []string {
	allImages := make([]string, 0)
	for _, file := range files {
		if file != nil {
			var beforePath string
			if fileType == IsFileType {
				beforePath = config.GetConfig().StaticFilePath
			} else {
				beforePath = config.GetConfig().StaticAvatarPath
			}

			ext := filepath.Ext(file.Filename)
			if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
				log.Printf("图片格式不支持")
				continue
			}
			if file.Size > 1920*1080*10 {
				log.Printf("图片大小超过限制")
				continue
			}
			filename := GenerateUUID() + ext
			file.Filename = filename
			filePath := beforePath + "/" + filename
			// 打开上传的文件
			src, err := file.Open()
			if err != nil {
				continue
			}
			defer src.Close()

			// 创建目标文件
			dst, err := os.Create(filePath)
			if err != nil {
				continue
			}
			defer dst.Close()
			// 保存文件
			if _, err := io.Copy(dst, src); err != nil {
				continue
			}
			allImages = append(allImages, filename)
		}
	}
	return allImages
}
