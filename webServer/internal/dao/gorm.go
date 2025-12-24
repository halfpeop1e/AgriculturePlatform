package dao

import (
	"fmt"
	"go-agriculture/internal/config"
	"go-agriculture/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func init() {
	conf := config.GetConfig()
	user := conf.PostgreConfig.User
	password := conf.PostgreConfig.Password
	host := conf.PostgreConfig.Host
	port := conf.PostgreConfig.Port
	dbName := conf.PostgreConfig.Db
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName)
	log.Println(dsn)
	// dsn := fmt.Sprintf("%s@unix(/var/run/mysqld/mysqld.sock)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, appName)
	var err error
	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	GormDB.AutoMigrate(&model.User{}, &model.Knowledge{}, &model.Order{}, &model.Product{},
		&model.LoanProduct{}, &model.Loan{}, &model.Expert{}, &model.ExpertBooking{}, model.ExpertCase{},
		&model.Question{}, &model.Answer{}, &model.Post{}, &model.PostLike{}, &model.PostReply{}, &model.PostReplyLike{}) // 自动迁移，如果没有建表，会自动创建对应的表
}
