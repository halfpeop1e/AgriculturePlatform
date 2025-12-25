package model

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	Id        int64          `gorm:"column:id;primarykey;comment:自增id"`
	Uuid      string         `gorm:"column:uuid;uniqueIndex;type:varchar(20);not null;comment:贷款uuid"`
	UserId    string         `gorm:"column:user_id;index;type:varchar(20);not null;comment:用户名ID"`
	ProductID string         `gorm:"column:product_id;type:varchar(20);not null;comment:贷款金额"`
	Amount    float64        `gorm:"column:amount;type:decimal(10,2);not null;comment:贷款金额"`
	Deadline  int64          `gorm:"column:deadline;type:int;not null;comment:贷款期限"`
	PayLine   int64          `gorm:"column:pay_line;type:int;not null;default:-1;comment:还款期限"`
	Status    int16          `gorm:"column:phone;type:int;not null;comment:贷款状态 0-申请中 1-已通过 2-已拒绝 3-已还款"`
	CreatAt   time.Time      `gorm:"column:create_at;type:timestamp;not null;comment:创建时间"`
	DeleteAt  gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}
