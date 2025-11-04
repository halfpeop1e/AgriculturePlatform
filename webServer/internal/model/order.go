package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id         int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid       string         `gorm:"column:uuid;uniqueIndex;type:varchar(20);not null;comment:订单uuid"`
	name       string         `gorm:"column:name;index;type:varchar(20);not null;comment:用户名"`
	quantity   int64          `gorm:"column:quantity;type:decimal(10,2);not null;comment:数量"`
	totalprice float64        `gorm:"column:totalprice;type:decimal(10,2);not null;comment:总价"`
	status     string         `gorm:"column:status;type:varchar(20);not null;comment:状态"`
	Ordertype  string         `gorm:"column:ordertype;type:varchar(20);not null;comment:类型"`
	CreatAt    time.Time      `gorm:"column:create_at;type:datetime;not null;comment:创建时间"`
	DeleteAt   gorm.DeletedAt `gorm:"column:delete_at;type:datetime;comment:删除时间"`
}
