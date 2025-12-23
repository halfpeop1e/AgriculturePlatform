package model

import (
	"time"

	"gorm.io/gorm"
)

// Post 帖子表
type Post struct {
	Id           int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid         string         `gorm:"column:uuid;uniqueIndex;type:varchar(36);not null;comment:帖子uuid"`
	Title        string         `gorm:"column:title;index;type:varchar(200);not null;comment:帖子标题"`
	Content      string         `gorm:"column:content;type:text;not null;comment:帖子内容"`
	AuthorId     string         `gorm:"column:author_id;index;type:varchar(36);not null;comment:作者ID"`
	AuthorName   string         `gorm:"column:author_name;type:varchar(50);not null;comment:作者姓名"`
	AuthorAvatar string         `gorm:"column:author_avatar;type:varchar(255);comment:作者头像"`
	LikeCount    int            `gorm:"column:like_count;type:int;default:0;not null;comment:点赞数"`
	ReplyCount   int            `gorm:"column:reply_count;type:int;default:0;not null;comment:回复数"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;not null;comment:创建时间"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;not null;comment:更新时间"`
	DeleteAt     gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}

// PostReply 帖子回复表
type PostReply struct {
	Id           int64          `gorm:"column:id;primarykey;comment:自增id"`
	PostId       string         `gorm:"column:post_id;index;type:varchar(36);not null;comment:帖子ID"`
	Content      string         `gorm:"column:content;type:text;not null;comment:回复内容"`
	AuthorId     string         `gorm:"column:author_id;index;type:varchar(36);not null;comment:作者ID"`
	AuthorName   string         `gorm:"column:author_name;type:varchar(50);not null;comment:作者姓名"`
	AuthorAvatar string         `gorm:"column:author_avatar;type:varchar(255);comment:作者头像"`
	LikeCount    int            `gorm:"column:like_count;type:int;default:0;not null;comment:点赞数"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;not null;comment:创建时间"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;not null;comment:更新时间"`
	DeleteAt     gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}

// PostLike 帖子点赞表（用户对帖子的点赞记录）
type PostLike struct {
	Id        int64     `gorm:"column:id;primarykey;comment:自增id"`
	PostId    string    `gorm:"column:post_id;index;type:varchar(36);not null;comment:帖子ID"`
	UserId    string    `gorm:"column:user_id;index;type:varchar(36);not null;comment:用户ID"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;comment:创建时间"`
}

// PostReplyLike 帖子回复点赞表（用户对回复的点赞记录）
type PostReplyLike struct {
	Id        int64     `gorm:"column:id;primarykey;comment:自增id"`
	ReplyId   int64     `gorm:"column:reply_id;index;type:int;not null;comment:回复ID"`
	UserId    string    `gorm:"column:user_id;index;type:varchar(36);not null;comment:用户ID"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;comment:创建时间"`
}
