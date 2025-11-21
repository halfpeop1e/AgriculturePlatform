package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	// 测试生成UUID
	uuid := GenerateUUID()

	assert.NotEmpty(t, uuid)
	assert.Equal(t, 20, len(uuid))
}

func TestGenerateUUIDUniqueness(t *testing.T) {
	// 测试生成的UUID唯一性
	uuids := make(map[string]bool)

	// 生成100个UUID验证唯一性
	for i := 0; i < 100; i++ {
		uuid := GenerateUUID()
		assert.False(t, uuids[uuid], "UUID重复: %s", uuid)
		uuids[uuid] = true
		assert.Equal(t, 20, len(uuid))
	}
}

func TestGenerateUUIDFormat(t *testing.T) {
	// 测试UUID格式
	uuid := GenerateUUID()

	// 验证只包含字母数字
	for _, char := range uuid {
		assert.True(t, (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z'),
			"UUID包含非法字符: %c", char)
	}
}

func TestGenerateMultipleUUIDs(t *testing.T) {
	// 测试多次生成UUID
	uuid1 := GenerateUUID()
	uuid2 := GenerateUUID()
	uuid3 := GenerateUUID()

	assert.NotEqual(t, uuid1, uuid2)
	assert.NotEqual(t, uuid2, uuid3)
	assert.NotEqual(t, uuid1, uuid3)

	assert.Equal(t, 20, len(uuid1))
	assert.Equal(t, 20, len(uuid2))
	assert.Equal(t, 20, len(uuid3))
}
