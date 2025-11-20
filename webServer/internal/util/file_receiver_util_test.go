package util

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileReceiverUtil(t *testing.T) {
	// 创建测试用的multipart文件
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 创建测试图片数据
	testData := []byte("fake image data for testing")
	part, err := writer.CreateFormFile("file", "test.jpg")
	require.NoError(t, err)
	_, err = part.Write(testData)
	require.NoError(t, err)

	writer.Close()

	// 创建HTTP请求
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 解析multipart form
	err = req.ParseMultipartForm(32 << 20) // 32MB
	require.NoError(t, err)

	// 获取文件
	files := req.MultipartForm.File["file"]

	// 测试文件接收
	result := FileReceiverUtil(files, IsFileType)

	assert.Len(t, result, 1)
	assert.NotEmpty(t, result[0])
	assert.True(t, len(result[0]) > 4) // UUID + extension

	// 清理测试文件
	if len(result) > 0 {
		// 这里应该根据实际配置获取文件路径进行清理
		_ = os.Remove(result[0])
	}
}

func TestFileReceiverUtilWithInvalidExtension(t *testing.T) {
	// 创建无效扩展名的文件
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	testData := []byte("fake file data")
	part, err := writer.CreateFormFile("file", "test.txt")
	require.NoError(t, err)
	_, err = part.Write(testData)
	require.NoError(t, err)

	writer.Close()

	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	err = req.ParseMultipartForm(32 << 20)
	require.NoError(t, err)

	files := req.MultipartForm.File["file"]

	// 测试无效扩展名应该被跳过
	result := FileReceiverUtil(files, IsFileType)
	assert.Empty(t, result)
}

func TestFileReceiverUtilWithNilFile(t *testing.T) {
	// 测试nil文件
	var files []*multipart.FileHeader
	result := FileReceiverUtil(files, IsFileType)
	assert.Empty(t, result)
}

func TestFileReceiverUtilWithAvatarType(t *testing.T) {
	// 创建测试用的multipart文件
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	testData := []byte("fake avatar data for testing")
	part, err := writer.CreateFormFile("file", "avatar.png")
	require.NoError(t, err)
	_, err = part.Write(testData)
	require.NoError(t, err)

	writer.Close()

	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	err = req.ParseMultipartForm(32 << 20)
	require.NoError(t, err)

	files := req.MultipartForm.File["file"]

	// 测试头像文件接收
	result := FileReceiverUtil(files, IsAvatarType)

	assert.Len(t, result, 1)
	assert.NotEmpty(t, result[0])
}
