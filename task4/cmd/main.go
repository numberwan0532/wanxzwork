package main

import (
	"log"

	"github.com/numberwan0532/wanxzwork/task4/configs"
	"github.com/numberwan0532/wanxzwork/task4/internal/api"
	"github.com/numberwan0532/wanxzwork/task4/internal/dao"
	"github.com/numberwan0532/wanxzwork/task4/internal/model"
	"github.com/numberwan0532/wanxzwork/task4/pkg/logs"
)

func main() {
	appLog := logs.InitLogger()
	//加载配置信息
	config, err := configs.InitConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	//连接数据库
	dao.InitDB(config, appLog)
	//创建数据表
	dao.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Commnet{})
	//启动web
	api.Start(config, appLog)
}
