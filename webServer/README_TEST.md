# Go 单元测试指南

本文档说明了如何运行和理解为 Go 农业平台项目编写的单元测试。

## 测试文件结构

```
webServer/
├── internal/util/
│   ├── jwt_util_test.go              # JWT 工具函数测试
│   ├── uuid_util_test.go             # UUID 生成工具测试
│   ├── file_receiver_util_test.go    # 文件接收工具测试
│   ├── file_receiver_base64_util_test.go # Base64文件接收测试
│   └── file_delete_util_test.go      # 文件删除工具测试
├── internal/server/
│   └── user_info_server_test.go      # 用户信息服务测试
└── go_test.go                        # 主测试文件和集成测试
```

## 运行测试

### 运行所有测试
```bash
cd webServer
go test ./...
```

### 运行特定包的测试
```bash
# 运行工具包测试
go test ./internal/util

# 运行服务器测试
go test ./internal/server
```

### 运行特定测试函数
```bash
# 运行JWT测试
go test -run TestGenerateToken ./internal/util

# 运行用户登录测试
go test -run TestUserInfoService_Login ./internal/server
```

### 显示测试覆盖率
```bash
# 显示覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 详细测试输出
```bash
# 显示详细输出
go test -v ./...

# 显示测试执行时间
go test -v -run=TestGenerateToken ./internal/util
```

## 测试内容说明

### 1. JWT 工具测试 (`jwt_util_test.go`)
- `TestGenerateToken`: 测试JWT token生成
- `TestParseToken`: 测试JWT token解析
- `TestParseInvalidToken`: 测试无效token解析
- `TestTokenExpiration`: 测试token过期时间
- `TestJWTClaimsStructure`: 测试JWT声明结构

### 2. UUID 工具测试 (`uuid_util_test.go`)
- `TestGenerateUUID`: 测试UUID生成基本功能
- `TestGenerateUUIDUniqueness`: 测试UUID唯一性
- `TestGenerateUUIDFormat`: 测试UUID格式
- `TestGenerateMultipleUUIDs`: 测试多次生成UUID的唯一性

### 3. 文件处理测试
- `file_receiver_util_test.go`: 测试multipart文件接收
- `file_receiver_base64_util_test.go`: 测试Base64文件接收
- `file_delete_util_test.go`: 测试文件删除功能

### 4. 用户服务测试 (`user_info_server_test.go`)
- `TestUserInfoService_Login`: 测试用户登录功能
- `TestUserInfoService_Register`: 测试用户注册功能
- `TestUserInfoService_GetUserInfo`: 测试获取用户信息功能

## 测试依赖

测试使用了以下依赖：
- `github.com/stretchr/testify`: 断言库
- `gorm.io/driver/sqlite`: 内存SQLite数据库（用于服务层测试）

## 运行测试前的准备

1. 确保已安装所有依赖：
```bash
go mod download
```

2. 确保测试环境配置正确

## 测试最佳实践

1. **测试命名**: 使用 `Test` 前缀，描述性命名
2. **测试隔离**: 每个测试独立运行，不依赖其他测试
3. **断言**: 使用 testify/assert 进行清晰的断言
4. **模拟**: 使用内存数据库进行服务层测试
5. **清理**: 测试后清理临时数据和文件

## 持续集成

这些测试可以集成到CI/CD流水线中：
```yaml
# GitHub Actions 示例
- name: Run tests
  run: |
    cd webServer
    go test -v -race -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html
```

## 添加新测试

当添加新功能时，请遵循以下步骤：
1. 为新功能编写对应的测试文件
2. 确保测试覆盖正常路径和错误路径
3. 运行测试确保通过
4. 更新本文档说明新的测试内容