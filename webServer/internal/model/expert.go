package model

import (
	"time"

	"gorm.io/gorm"
)

type Expert struct {
	Id            int64          `gorm:"column:id;primarykey;comment:自增id"`
	Name          string         `gorm:"column:name;index;type:varchar(50);not null;comment:专家姓名"`
	Avatar        string         `gorm:"column:avatar;type:varchar(255);default:'default_expert.png';comment:专家头像"`
	Field         string         `gorm:"column:field;index;type:varchar(100);not null;comment:专业领域"`
	Introduction  string         `gorm:"column:introduction;type:text;comment:专家简介"`
	Skills        string         `gorm:"column:skills;type:json;default:'[]';comment:擅长技能"`
	Education     string         `gorm:"column:education;type:text;comment:教育背景"`
	Experience    string         `gorm:"column:experience;type:text;comment:工作经历"`
	Certification string         `gorm:"column:certification;type:json;default:'[]';comment:资质证书"`
	ConsultCount  int            `gorm:"column:consult_count;type:int;default:0;not null;comment:咨询次数"`
	Rating        float64        `gorm:"column:rating;type:decimal(3,2);default:5.00;not null;comment:好评率(0-5)"`
	ResponseTime  string         `gorm:"column:response_time;type:varchar(50);comment:平均响应时间"`
	AvailableTime string         `gorm:"column:available_time;type:json;default:'[]';comment:可咨询时间"`
	ServiceScope  string         `gorm:"column:service_scope;type:json;default:'[]';comment:服务范围"`
	Price         float64        `gorm:"column:price;type:decimal(10,2);comment:咨询费用"`
	Status        string         `gorm:"column:status;type:varchar(20);default:'active';not null;comment:专家状态(active/inactive)"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;not null;comment:创建时间"`
	DeleteAt      gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}

type ExpertCase struct {
	Id        int64          `gorm:"column:id;primarykey;comment:自增id"`
	ExpertId  string         `gorm:"column:expert_id;index;type:varchar(36);not null;comment:专家ID"`
	Question  string         `gorm:"column:question;type:text;not null;comment:咨询问题"`
	Answer    string         `gorm:"column:answer;type:text;comment:专家回答"`
	CaseDate  time.Time      `gorm:"column:case_date;type:timestamp;not null;comment:案例日期"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null;comment:创建时间"`
	DeleteAt  gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}

// ExpertBooking 专家咨询预约记录表
type ExpertBooking struct {
	Id          int64          `gorm:"column:id;primarykey;comment:自增id"`
	ExpertId    string         `gorm:"column:expert_id;index;type:varchar(36);not null;comment:专家ID"`
	UserId      string         `gorm:"column:user_id;index;type:varchar(36);not null;comment:用户ID"`
	UserName    string         `gorm:"column:user_name;type:varchar(50);not null;comment:用户姓名"`
	Question    string         `gorm:"column:question;type:text;not null;comment:咨询问题"`
	ContactInfo string         `gorm:"column:contact_info;type:varchar(100);comment:联系方式"`
	BookDate    string         `gorm:"column:book_date;type:varchar(50);not null;comment:预约时间"`
	Status      string         `gorm:"column:status;type:varchar(20);default:'pending';not null;comment:预约状态(pending/confirmed/completed/cancelled)"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp;not null;comment:创建时间"`
	DeleteAt    gorm.DeletedAt `gorm:"column:delete_at;type:timestamp;comment:删除时间"`
}
