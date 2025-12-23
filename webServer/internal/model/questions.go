package model

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	Id int64 `gorm:"column:id;primarykey;comment:自增id"`

	ExpertId string `gorm:"column:expert_id;index;type:varchar(30);not null;comment:专家ID"`

	Title string `gorm:"column:title;index;type:varchar(200);not null;comment:问题标题"`

	Content string `gorm:"column:content;type:text;comment:问题内容"`

	Tags []string `gorm:"column:tags;type:json;default:'[]';comment:问题标签"`

	Author string `gorm:"column:author;type:varchar(45);not null;comment:问题作者"`

	AuthorId string `gorm:"column:author_id;index;type:varchar(36);not null;comment:作者ID"`

	Date time.Time `gorm:"column:date;type:timestamp;not null;comment:更新时间"`

	AnswerCount int `gorm:"column:answer_count;type:int;default:0;not null;comment:答案数量"`

	IsAnswered bool `gorm:"column:is_answered;type:boolean;default:false;not null;comment:是否已回答"`

	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}
