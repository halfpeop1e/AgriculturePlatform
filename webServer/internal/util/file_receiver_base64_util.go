package util

import (
	"encoding/base64"
	"errors"
	"go-agriculture/internal/config"
	"log"
	"os"
	"strings"
)

func FileReceiverBase64Util(files []string, fileType FileType) ([]string, error) {
	allImages := make([]string, 0)
	var beforePath string
	if fileType == IsFileType {
		beforePath = config.GetConfig().StaticFilePath
	} else {
		beforePath = config.GetConfig().StaticAvatarPath
	}
	for _, file := range files {
		if file != "" {
			if len(file) > 1920*1080*2 {
				log.Printf("图片大小超过限制")
				continue
			}
			parts := strings.Split(file, ",")
			if len(parts) != 2 {
				return nil, errors.New("格式错误")
			}
			// 检查格式是否为 Base64 图片
			if !strings.HasPrefix(file, "data:image/") {
				return nil, errors.New("avatar格式错误，请提供有效的base64图片")
			}
			// 解码 Base64
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				return nil, errors.New("解码失败")
			}

			var ext string
			if strings.Contains(parts[0], "image/jpeg") {
				ext = ".jpg"
			} else if strings.Contains(parts[0], "image/png") {
				ext = ".png"
			} else if strings.Contains(parts[0], "image/gif") {
				ext = ".gif"
			} else if strings.Contains(parts[0], "image/webp") {
				ext = ".webp"
			} else {
				return nil, errors.New("不支持的图片格式")
			}
			filename := GenerateUUID() + ext
			filePath := beforePath + "/" + filename

			// 创建目标文件
			dst, err := os.Create(filePath)
			if err != nil {
				continue
			}
			defer dst.Close()
			// 写入图片数据
			if _, err := dst.Write(decoded); err != nil {
				continue
			}
			allImages = append(allImages, filename)
		}
	}
	return allImages, nil
}
