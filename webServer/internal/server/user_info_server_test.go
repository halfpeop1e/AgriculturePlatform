package gorm

import (
	"testing"
	"time"

	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 设置测试数据库
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	// 自动迁移
	err = db.AutoMigrate(&model.User{})
	require.NoError(t, err)

	return db
}

// 清理测试数据库
func cleanupTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM user_info")
}

func TestUserInfoService_Login(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)

	// 临时替换dao.GormDB
	originalDB := dao.GormDB
	dao.GormDB = db
	defer func() { dao.GormDB = originalDB }()

	// 创建测试用户
	testUser := model.User{
		Uuid:     "test123",
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password123",
		CreatAt:  time.Now(),
	}
	err := db.Create(&testUser).Error
	require.NoError(t, err)

	// 测试登录成功
	loginReq := request.LoginRequest{
		Username: "test@example.com",
		Password: "password123",
	}

	message, resp, code := UserInfoService.Login(loginReq)

	assert.Equal(t, "登陆成功", message)
	assert.NotNil(t, resp)
	assert.Equal(t, 0, code)
	assert.Equal(t, "test123", resp.UserId)
	assert.NotEmpty(t, resp.Tokens)
}

func TestUserInfoService_LoginWithInvalidCredentials(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)

	originalDB := dao.GormDB
	dao.GormDB = db
	defer func() { dao.GormDB = originalDB }()

	// 测试用户不存在
	loginReq := request.LoginRequest{
		Username: "nonexistent@example.com",
		Password: "password123",
	}

	message, resp, code := UserInfoService.Login(loginReq)

	assert.Equal(t, "用户不存在", message)
	assert.Nil(t, resp)
	assert.Equal(t, -1, code)
}

func TestUserInfoService_LoginWithWrongPassword(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)

	originalDB := dao.GormDB
	dao.GormDB = db
	defer func() { dao.GormDB = originalDB }()

	// 创建测试用户
	testUser := model.User{
		Uuid:     "test123",
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password123",
		CreatAt:  time.Now(),
	}
	err := db.Create(&testUser).Error
	require.NoError(t, err)

	// 测试密码错误
	loginReq := request.LoginRequest{
		Username: "test@example.com",
		Password: "wrongpassword",
	}

	message, resp, code := UserInfoService.Login(loginReq)

	assert.Equal(t, "密码错误", message)
	assert.Nil(t, resp)
	assert.Equal(t, -1, code)
}

func TestUserInfoService_Register(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)

	originalDB := dao.GormDB
	dao.GormDB = db
	defer func() { dao.GormDB = originalDB }()

	// 测试注册成功
	registerReq := request.RegisterRequest{
		Username:    "newuser",
		Password:    "password123",
		ComPassword: "password123",
		Email:       "newuser@example.com",
	}

	message, resp, code := UserInfoService.Register(registerReq)

	assert.Equal(t, "注册成功", message)
	assert.Equal(t, 0, code)
	assert.Nil(t, resp) // 注册成功时不返回token

	// 验证用户已创建
	var user model.User
	err := db.Where("email = ?", "newuser@example.com").First(&user).Error
	assert.NoError(t, err)
	assert.Equal(t, "newuser", user.Name)
	assert.Equal(t, "newuser@example.com", user.Email)
}

func TestUserInfoService_RegisterWithExistingEmail(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)

	originalDB := dao.GormDB
	dao.GormDB = db
	defer func() { dao.GormDB = originalDB }()

	// 创建现有用户
	testUser := model.User{
		Uuid:     "existing123",
		Name:     "existinguser",
		Email:    "existing@example.com",
		Password: "password123",
		CreatAt:  time.Now(),
	}
	err := db.Create(&testUser).Error
	require.NoError(t, err)

	// 测试重复注册
	registerReq := request.RegisterRequest{
		Username:    "newuser",
		Password:    "password123",
		ComPassword: "password123",
		Email:       "existing@example.com", // 已存在的邮箱
	}

	message, resp, code := UserInfoService.Register(registerReq)

	assert.Equal(t, "该用户已经存在，注册失败", message)
	assert.Equal(t, -2, code)
	assert.Nil(t, resp)
}

func TestUserInfoService_RegisterWithPasswordMismatch(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)

	originalDB := dao.GormDB
	dao.GormDB = db
	defer func() { dao.GormDB = originalDB }()

	// 测试密码不匹配
	registerReq := request.RegisterRequest{
		Username:    "newuser",
		Password:    "password123",
		ComPassword: "different123",
		Email:       "newuser@example.com",
	}

	message, resp, code := UserInfoService.Register(registerReq)

	assert.Equal(t, "两次密码不一致", message)
	assert.Equal(t, -1, code)
	assert.Nil(t, resp)
}

func TestUserInfoService_GetUserInfo(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)

	originalDB := dao.GormDB
	dao.GormDB = db
	defer func() { dao.GormDB = originalDB }()

	// 创建测试用户
	testUser := model.User{
		Uuid:       "test123",
		Name:       "testuser",
		Email:      "test@example.com",
		Bio:        "Test bio",
		Phone:      "1234567890",
		Adress:     "Test address",
		Location:   "Test location",
		Avatar:     "test.jpg",
		Tags:       `["tag1", "tag2"]`,
		CreatAt:    time.Now(),
		ActiveTime: time.Now(),
	}
	err := db.Create(&testUser).Error
	require.NoError(t, err)

	// 测试获取用户信息
	message, resp, code := UserInfoService.GetUserInfo("test123")

	assert.Equal(t, "获取用户信息成功", message)
	assert.Equal(t, 0, code)
	assert.NotNil(t, resp)
	assert.Equal(t, "testuser", resp.Nickname)
	assert.Equal(t, "test@example.com", resp.Email)
	assert.Equal(t, "Test bio", resp.Bio)
	assert.Equal(t, "1234567890", resp.Phone)
	assert.Equal(t, "Test address", resp.Address)
	assert.Equal(t, "Test location", resp.Location)
	assert.Contains(t, resp.Tags, "tag1")
	assert.Contains(t, resp.Tags, "tag2")
}

func TestUserInfoService_GetUserInfoWithNonExistentUser(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(db)

	originalDB := dao.GormDB
	dao.GormDB = db
	defer func() { dao.GormDB = originalDB }()

	// 测试获取不存在用户的信息
	message, resp, code := UserInfoService.GetUserInfo("nonexistent")

	assert.Equal(t, "用户不存在", message)
	assert.Equal(t, -1, code)
	assert.Nil(t, resp)
}
