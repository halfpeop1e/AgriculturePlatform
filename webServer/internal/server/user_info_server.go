package gorm

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-agriculture/internal/config"
	"go-agriculture/internal/dao"
	"go-agriculture/internal/dto/request"
	"go-agriculture/internal/dto/respond"
	"go-agriculture/internal/model"
	"go-agriculture/internal/util"
	"log"
	"time"

	"gorm.io/gorm"
)

type userInfoService struct {
}

var UserInfoService = &userInfoService{}

func (u *userInfoService) Login(loginReq request.LoginRequest) (string, *respond.LoginRespond, int) {
	if len(loginReq.Username) == 0 || len(loginReq.Password) == 0 {
		return "用户名或密码不能为空", nil, -1
	}
	if len(loginReq.Username) > 50 || len(loginReq.Password) > 50 {
		return "用户名或密码长度超过限制", nil, -1
	}
	password := loginReq.Password
	var user model.User
	res := dao.GormDB.Where("email = ?", loginReq.Username).First(&user)
	if res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return "用户不存在", nil, -1
		}
		return "查询用户信息失败", nil, -1
	}
	if user.Password != password {
		return "密码错误", nil, -1
	}
	user.ActiveTime = time.Now()
	dao.GormDB.Save(&user)
	token, err := util.GenerateToken(user.Uuid)
	if err != nil {
		return "生成token失败", nil, -1
	}
	loginResp := &respond.LoginRespond{
		UserId: user.Uuid,
		Tokens: token,
	}
	return "登陆成功", loginResp, 0
}

func (u *userInfoService) checkNameExist(name string, email string) (string, int) {
	var user model.User
	// gorm默认排除软删除，所以翻译过来的select语句是SELECT * FROM `user_info` WHERE telephone = '18089596095' AND `user_info`.`deleted_at` IS NULL ORDER BY `user_info`.`id` LIMIT 1
	if res := dao.GormDB.Where("name = ? OR email = ?", name, email).First(&user); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return "", 0
		}
		return "查询失败", -1
	}
	return "该用户已经存在，注册失败", -2
}

func (u *userInfoService) Register(registerReq request.RegisterRequest) (string, *respond.LoginRespond, int) {
	if len(registerReq.Username) == 0 || len(registerReq.Password) == 0 {
		return "用户名或密码不能为空", nil, -1
	}
	if len(registerReq.Username) > 20 || len(registerReq.Password) > 20 || len(registerReq.Email) > 50 {
		return "长度超过限制", nil, -1
	}
	if registerReq.Password != registerReq.ComPassword {

		return "两次密码不一致", nil, -1

	}
	if !registerReq.Isverified {
		return "请先验证邮箱", nil, -1
	}

	var user model.User
	message, ret := u.checkNameExist(registerReq.Username, registerReq.Email)
	if ret != 0 {
		return message, nil, ret
	}
	user.Name = registerReq.Username
	user.Password = registerReq.Password
	user.Email = registerReq.Email
	user.Uuid = util.GenerateUUID()
	user.CreatAt = time.Now()
	user.Avatar = "default.png"
	if res := dao.GormDB.Create(&user); res.Error != nil {
		return "注册失败", nil, -1
	}
	return "注册成功", nil, 0
}

func (u *userInfoService) GetUserInfo(userId string) (string, *respond.UserInfoRespond, int) {
	var user model.User
	res := dao.GormDB.Where("uuid = ?", userId).First(&user)
	if res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return "用户不存在", nil, -1
		}
		return "查询用户信息失败", nil, -1
	}
	var jsonTags []string
	if user.Tags == "" {
		jsonTags = []string{}
	} else {
		if err := json.Unmarshal([]byte(user.Tags), &jsonTags); err != nil {
			log.Printf("json转换失败: %v", err)
			return "json转换失败", nil, -1
		}
	}

	avatarUrl := "http://" + fmt.Sprint(config.GetConfig().MainConfig.ServerIP) + ":" + fmt.Sprint(config.GetConfig().MainConfig.Port) + "/static/avatars/" + user.Avatar
	userInfoResp := &respond.UserInfoRespond{
		Nickname:   user.Name,
		Email:      user.Email,
		Avatar:     avatarUrl,
		Bio:        user.Bio,
		Tags:       jsonTags,
		Location:   user.Location,
		JoinedAt:   user.CreatAt.Format("2006年01月02日 15:04:05"),
		Phone:      user.Phone,
		Address:    user.Adress,
		LastActive: user.ActiveTime.Format("2006年01月02日 15:04:05"),
	}
	return "获取用户信息成功", userInfoResp, 0
}
