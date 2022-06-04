package main

import (
	"context"
	"fmt"
	"gopl.io/dal/model"
	"gopl.io/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/go-gateway?charset=utf8mb4&parseTime=True&loc=Local"), gormConfig())

	admin := model.GatewayAdmin{
		UserName:   "admin",
		Password:   "123456",
		Salt:       "zzx",
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	adminDao := query.Use(db).GatewayAdmin

	err := adminDao.WithContext(context.Background()).Create(&admin)
	if err != nil {
		fmt.Println(err)
	}
}

// gormConfig 根据配置决定是否开启日志
func gormConfig() *gorm.Config {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,            // Log level
			Colorful:      true,                   // 禁用彩色打印
		},
	)
	var config = &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: newLogger,
	}
	return config
}
