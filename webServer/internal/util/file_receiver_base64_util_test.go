package util

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileReceiverBase64Util(t *testing.T) {
	// 创建测试用的base64图片数据
	testImageData := []byte("fake image data for testing")
	encoded := base64.StdEncoding.EncodeToString(testImageData)
	base64Image := "data:image/jpeg;base64," + encoded
	
	files := []string{base64Image}
	
	result, err := FileReceiverBase64Util(files, IsFileType)
	
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.NotEmpty(t, result[0])
	assert.True(t, strings.HasSuffix(result[0], ".jpg"))
}

func TestFileReceiverBase64UtilWithPNG(t *testing.T) {
	// 测试PNG格式
	testImageData := []byte("fake png data for testing")
	encoded := base64.StdEncoding.EncodeToString(testImageData)
	base64Image := "data:image/png;base64," + encoded
	
	files := []string{base64Image}
	
	result, err := FileReceiverBase64Util(files, IsFileType)
	
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.True(t, strings.HasSuffix(result[0], ".png"))
}

func TestFileReceiverBase64UtilWithInvalidFormat(t *testing.T) {
	// 测试无效格式
	files := []string{"invalid_base64_data"}
	
	result, err := FileReceiverBase64Util(files, IsFileType)
	
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestFileReceiverBase64UtilWithUnsupportedFormat(t *testing.T) {
	// 测试不支持的图片格式
	testImageData := []byte("fake bmp data")
	encoded := base64.StdEncoding.EncodeToString(testImageData)
	base64Image := "data:image/bmp;base64," + encoded
	
	files := []string{base64Image}
	
	result, err := FileReceiverBase64Util(files, IsFileType)
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "不支持的图片格式")
	assert.Nil(t, result)
}

func TestFileReceiverBase64UtilWithEmptyString(t *testing.T) {
	// 测试空字符串
	files := []string{""}
	
	result, err := FileReceiverBase64Util(files, IsFileType)
	
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestFileReceiverBase64UtilWithMultipleFiles(t *testing.T) {
	// 测试多个文件
	testImageData1 := []byte("fake image data 1")
	testImageData2 := []byte("fake image data 2")
	
	encoded1 := base64.StdEncoding.EncodeToString(testImageData1)
	encoded2 := base64.StdEncoding.EncodeToString(testImageData2)
	
	base64Image1 := "data:image/jpeg;base64," + encoded1
	base64Image2 := "data:image/png;base64," + encoded2
	
	files := []string{base64Image1, base64Image2, ""} // 包含一个空字符串
	
	result, err := FileReceiverBase64Util(files, IsFileType)
	
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.True(t, strings.HasSuffix(result[0], ".jpg"))
	assert.True(t, strings.HasSuffix(result[1], ".png"))
}

func TestFileReceiverBase64UtilWithAvatarType(t *testing.T) {
	// 测试头像类型
	testImageData := []byte("fake avatar data")
	encoded := base64.StdEncoding.EncodeToString(testImageData)
	base64Image := "data:image/webp;base64," + encoded
	
	files := []string{base64Image}
	
	result, err := FileReceiverBase64Util(files, IsAvatarType)
	
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.True(t, strings.HasSuffix(result[0], ".webp"))
}