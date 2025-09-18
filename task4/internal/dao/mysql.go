package dao

import (
	"fmt"
	"time"

	"github.com/numberwan0532/wanxzwork/task4/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config *configs.Config, appLog *logrus.Logger) {
	// dsn := "root:123456abc@tcp(192.168.124.13:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName) // 构建DSN字符串

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	appLog.Info("数据库连接成功", DB)
	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed to connect pool")
	}
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// GlobalDB.AutoMigrate(&model.User{}, &model.Post{}, &model.Commnet{})
}
