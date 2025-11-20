package util

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// 测试生成token
	userID := "test123"
	token, err := GenerateToken(userID)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
	// 测试解析有效token
	userID := "test123"
	token, err := GenerateToken(userID)
	assert.NoError(t, err)

	claims, err := ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, "go-agriculture", claims.Issuer)
}

func TestParseInvalidToken(t *testing.T) {
	// 测试解析无效token
	invalidToken := "invalid.token.here"
	claims, err := ParseToken(invalidToken)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestTokenExpiration(t *testing.T) {
	// 测试token过期时间
	userID := "test123"
	token, err := GenerateToken(userID)
	assert.NoError(t, err)

	claims, err := ParseToken(token)
	assert.NoError(t, err)

	// 检查过期时间是否在未来
	assert.True(t, claims.ExpiresAt.After(time.Now()))
}

func TestJWTClaimsStructure(t *testing.T) {
	// 测试JWT声明结构
	userID := "test456"
	token, err := GenerateToken(userID)
	assert.NoError(t, err)

	// 手动解析token验证结构
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(*Claims)
	assert.True(t, ok)
	assert.Equal(t, userID, claims.UserID)
}
