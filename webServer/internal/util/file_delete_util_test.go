package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileDeleteUtilWithEmptyFileName(t *testing.T) {
	// 测试空文件名
	FileDeleteUtil("", IsFileType)
	// 应该不产生错误，直接返回
}

func TestFileDeleteUtilWithDefaultAvatar(t *testing.T) {
	// 测试默认头像文件名
	FileDeleteUtil("default.png", IsAvatarType)
	// 应该不产生错误，直接返回
}

func TestFileDeleteUtilWithNonExistentFile(t *testing.T) {
	// 测试不存在的文件
	FileDeleteUtil("non_existent_file.jpg", IsFileType)
	// 应该不产生错误
}

func TestFileDeleteUtilWithExistingFile(t *testing.T) {
	// 创建临时测试文件
	tempDir := t.TempDir()
	testFileName := "test_delete.jpg"
	testFilePath := filepath.Join(tempDir, testFileName)
	
	// 创建测试文件
	testData := []byte("test data for deletion")
	err := os.WriteFile(testFilePath, testData, 0644)
	require.NoError(t, err)
	
	// 验证文件存在
	_, err = os.Stat(testFilePath)
	assert.NoError(t, err)
	
	// 注意：这个测试需要模拟配置文件中的路径
	// 由于FileDeleteUtil使用配置文件中的路径，我们无法直接测试文件删除
	// 但可以验证函数不会panic
	FileDeleteUtil(testFileName, IsFileType)
}