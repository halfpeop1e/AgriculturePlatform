package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid       string         `gorm:"column:uuid;uniqueIndex;type:varchar(20);not null;comment:用户uuid"`
	Name       string         `gorm:"column:name;index;type:varchar(20);not null;comment:用户名"`
	Email      string         `gorm:"column:email;type:varchar(50);not null;comment:用户邮箱"`
	Phone      string         `gorm:"column:phone;type:varchar(20);comment:用户手机号"`
	Avatar     string         `gorm:"column:avatar;type:varchar(255);default:'default.png';not null;comment:用户头像"`
	Password   string         `gorm:"column:password;type:varchar(20);not null;comment:用户密码"`
	Tags       string         `gorm:"column:tags;type:json;default:'[]';comment:用户标签"`
	Location   string         `gorm:"column:location;type:varchar(50);comment:用户位置"`
	Adress     string         `gorm:"column:adress;type:varchar(50);comment:用户地址"`
	Bio        string         `gorm:"column:bio;type:varchar(255);default:'褪色者成为混沌之王，燃烧了一切，交界地陷入一片火海之中。梅琳娜表示自己将找到褪色者，给予他命定之死。';comment:用户简介"`
	ActiveTime time.Time      `gorm:"column:active_time;type:timestamp;comment:用户活跃时间"`
	CreatAt    time.Time      `gorm:"column:create_at;type:timestamp;not null;comment:创建时间"`
	DeleteAt   gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}
