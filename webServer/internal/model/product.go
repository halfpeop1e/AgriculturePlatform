package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid        string         `gorm:"column:uuid;uniqueIndex;type:varchar(36);not null;comment:商品uuid"`
	Name        string         `gorm:"column:name;index;type:varchar(200);not null;comment:商品名称"`
	Images      string         `gorm:"column:images;type:json;default:'[]';comment:商品图片列表"`
	Category    string         `gorm:"column:category;index;type:varchar(50);not null;comment:商品分类"`
	Description string         `gorm:"column:description;type:text;comment:商品描述"`
	Price       float64        `gorm:"column:price;type:decimal(10,2);not null;comment:商品价格"`
	Stock       int64          `gorm:"column:stock;type:int;default:0;not null;comment:商品库存"`
	Saler       string         `gorm:"column:saler;type:varchar(100);not null;comment:卖家名称"`
	SalerId     string         `gorm:"column:saler_id;index;type:varchar(36);not null;comment:卖家ID"`
	Status      string         `gorm:"column:status;type:int;default:1;not null;comment:商品状态0-下架1-上架"`
	CreatAt     time.Time      `gorm:"column:create_at;type:timestamp;not null;comment:创建时间"`
	Email       string         `gorm:"column:email;type:varchar(100);not null;comment:联系邮箱"`
	UpdateAt    time.Time      `gorm:"column:update_at;type:timestamp;not null;comment:更新时间"`
	DeleteAt    gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}
