package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id         int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid       string         `gorm:"column:uuid;uniqueIndex;type:varchar(20);not null;comment:订单uuid"`
	Name       string         `gorm:"column:name;index;type:varchar(20);not null;comment:用户名"`
	Quantity   int64          `gorm:"column:quantity;type:int;not null;comment:数量"`
	Totalprice float64        `gorm:"column:totalprice;type:decimal(10,2);not null;comment:总价"`
	Status     string         `gorm:"column:status;type:varchar(20);not null;comment:状态"`
	Ordertype  string         `gorm:"column:ordertype;type:varchar(20);not null;comment:类型"`
	ProductId  int64          `gorm:"column:product_id;type:int;not null;comment:产品id"`
	UserId     string         `gorm:"column:user_id;type:varchar(30);not null;comment:用户id"`
	Buyer      string         `gorm:"column:buyer;type:varchar(50);not null;comment:买家"`
	Saler      string         `gorm:"column:saler;type:varchar(50);not null;comment:卖家"`
	CreatAt    time.Time      `gorm:"column:create_at;type:timestamp;not null;comment:创建时间"`
	DeleteAt   gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}
