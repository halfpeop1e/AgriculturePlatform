package model

import (
	"time"

	"gorm.io/gorm"
)

type Knowledge struct {
	Id       int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid     string         `gorm:"column:uuid;uniqueIndex;type:varchar(20);not null;comment:知识uuid"`
	title    string         `gorm:"column:title;index;type:varchar(200);not null;comment:知识标题"`
	excerpt  string         `gorm:"column:excerpt;type:text;comment:知识摘要"`
	content  string         `gorm:"column:content;type:text;comment:知识内容"`
	image    string         `gorm:"column:image;type:varchar(255);default:'static/avatars/default.png';comment:知识图片"`
	tags     []string       `gorm:"column:tags;type:json;comment:知识标签"`
	author   string         `gorm:"column:author;type:varchar(20);not null;comment:知识作者"`
	UpdateAt time.Time      `gorm:"column:update_at;type:datetime;not null;comment:更新时间"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;type:datetime;comment:删除时间"`
}
