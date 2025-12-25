package model

import (
	"time"

	"gorm.io/gorm"
)

type AllOrder struct {
	Id         int64          `gorm:"column:id;primarykey;comment:自增id"`
	Year       int            `gorm:"column:year;type:int;not null;comment:年份"`
	Month      int            `gorm:"column:month;type:int;not null;comment:月份"`
	Amount     float64        `gorm:"column:amount;type:decimal(10,2);not null;comment:金额"`
	LoanName   string         `gorm:"column:loan_name;type:varchar(35);not null;comment:贷款名称"`
	LoanStatus string         `gorm:"column:loan_status;type:varchar(35);not null;comment:贷款状态"`
	LoanId     string         `gorm:"column:loan_id;type:varchar(35);not null;comment:贷款id"`
	CreatAt    time.Time      `gorm:"column:create_at;type:timestamp;not null;comment:创建时间"`
	DeleteAt   gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}
