package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id       int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid     string         `gorm:"column:uuid;uniqueIndex;type:varchar(20);not null;comment:用户uuid"`
	Name     string         `gorm:"column:name;index;type:varchar(20);not null;comment:用户名"`
	Email    string         `gorm:"column:email;type:varchar(20);not null;comment:用户邮箱"`
	Phone    string         `gorm:"column:phone;type:varchar(20);not null;comment:用户手机号"`
	Avatar   string         `gorm:"column:avatar;type:varchar(255);default:'static/avatars/default.png';not null;comment:用户头像"`
	Password string         `gorm:"column:password;type:varchar(20);not null;comment:用户密码"`
	CreatAt  time.Time      `gorm:"column:create_at;type:datetime;not null;comment:创建时间"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;type:datetime;comment:删除时间"`
}
