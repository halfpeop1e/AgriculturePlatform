package model

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	Id         int64          `gorm:"column:id;primarykey;comment:自增id"`
	QuestionId int64          `gorm:"column:question_id;index;not null;comment:问题ID"`
	Content    string         `gorm:"column:content;type:text;comment:答案内容"`
	Date       time.Time      `gorm:"column:date;type:timestamp;not null;comment:更新时间"`
	IsAccpet   bool           `gorm:"column:is_accepted;type:boolean;default:false;not null;comment:是否已回答"`
	DeleteAt   gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}
