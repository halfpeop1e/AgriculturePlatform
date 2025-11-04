package model

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	Id       int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid     string         `gorm:"column:uuid;uniqueIndex;type:varchar(20);not null;comment:贷款uuid"`
	UserId   string         `gorm:"column:name;index;type:varchar(20);not null;comment:用户名ID"`
	Amount   float64        `gorm:"column:email;type:decimal(10,2);not null;comment:贷款金额"`
	Reason   string         `gorm:"column:phone;type:varchar(20);not null;comment:贷款原因"`
	Status   int16          `gorm:"column:phone;type:int;not null;comment:贷款状态 0-申请中 1-已通过 2-已拒绝 3-已还款"`
	CreatAt  time.Time      `gorm:"column:create_at;type:datetime;not null;comment:创建时间"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;type:datetime;comment:删除时间"`
}
