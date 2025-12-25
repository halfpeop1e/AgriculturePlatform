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
	token, err := util.GenerateToken(user.Uuid, user.Role)
	if err != nil {
		return "生成token失败", nil, -1
	}
	loginResp := &respond.LoginRespond{
		UserId: user.Uuid,
		Role:   user.Role,
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
	user.Role = registerReq.Role
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

func (u *userInfoService) UpdateUserInfo(userId string, updateReq request.UpdateUserInfoRequest) (string, int) {
	if len(updateReq.Nickname) > 20 {
		return "昵称长度超过限制", -1
	}
	if len(updateReq.Bio) > 255 {
		return "简介长度超过限制", -1
	}
	if len(updateReq.Phone) > 20 {
		return "手机号长度超过限制", -1
	}
	if len(updateReq.Address) > 50 {
		return "地址长度超过限制", -1
	}
	if len(updateReq.Location) > 50 {
		return "地址长度超过限制", -1
	}

	var user model.User
	res := dao.GormDB.Where("uuid = ?", userId).First(&user)
	if res.Error != nil {
		log.Printf("Database error: %v", res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return "用户不存在", -1
		}
		return "查询用户信息失败", -1
	}
	if updateReq.Nickname != "" {
		user.Name = updateReq.Nickname
	}
	if updateReq.Bio != "" {
		user.Bio = updateReq.Bio
	}
	if updateReq.Avatar != "" {
		util.FileDeleteUtil(user.Avatar, util.IsAvatarType)
		files := make([]string, 1)
		files[0] = updateReq.Avatar
		newAvatar, err := util.FileReceiverBase64Util(files, util.IsAvatarType)
		if err != nil {
			return "文件上传失败", -1
		}
		user.Avatar = newAvatar[0]
	}
	if updateReq.Location != "" {
		user.Location = updateReq.Location
	}
	if updateReq.Phone != "" {
		user.Phone = updateReq.Phone
	}
	if updateReq.Address != "" {
		user.Adress = updateReq.Address
	}
	if updateReq.Tags != nil {
		tags, err := json.Marshal(updateReq.Tags)
		if err != nil {
			log.Printf("json转换失败: %v", err)
			return "json转换失败", -1
		}
		user.Tags = string(tags)
	}
	user.ActiveTime = time.Now()
	dao.GormDB.Save(&user)
	return "更新用户信息成功", 0
}

func (u *userInfoService) SafeUpdateInfo(userId string, req request.SafeUpdateUserInfoRequest) (string, int) {
	if len(req.Currentpassword) > 20 || len(req.Newpassword) > 20 {
		return "密码长度超过限制", -1
	}
	if len(req.Email) > 50 {
		return "邮箱长度错误", -1
	}
	var user model.User
	if res := dao.GormDB.Where("uuid = ?", userId).First(&user); res != nil {
		log.Printf("Database error: %v", res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return "用户不存在", -1
		}
		return "查询用户信息失败", -1
	}
	if user.Password != req.Currentpassword {
		return "旧密码错误", -1
	}
	if req.Newpassword == req.Currentpassword {
		return "两次密码一致", -1
	}
	if len(req.Newpassword) != 0 {
		user.Password = req.Newpassword
	}
	if len(req.Email) != 0 {
		user.Email = req.Email

	}
	dao.GormDB.Save(&user)
	return "更新用户信息成功", 0
}
