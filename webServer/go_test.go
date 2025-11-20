package main

import (
	"os"
	"testing"

	"go-agriculture/internal/util"
)

// TestMain 是测试的主入口点
func TestMain(m *testing.M) {
	// 在所有测试开始前执行的设置代码
	// 例如：初始化配置、连接测试数据库等
	
	// 运行测试
	code := m.Run()
	
	// 在所有测试结束后执行的清理代码
	os.Exit(code)
}

// 示例集成测试
func TestIntegrationExample(t *testing.T) {
	// 测试UUID生成
	uuid := util.GenerateUUID()
	if len(uuid) != 20 {
		t.Errorf("Expected UUID length 20, got %d", len(uuid))
	}
	
	// 测试JWT token生成和解析
	token, err := util.GenerateToken(uuid)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	
	claims, err := util.ParseToken(token)
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}
	
	if claims.UserID != uuid {
		t.Errorf("Expected user ID %s, got %s", uuid, claims.UserID)
	}
}