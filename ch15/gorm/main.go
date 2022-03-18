package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type SysUser struct {
	gorm.Model
	Username string    `json:"username"`
	RealName string    `json:"real_name" likeField:"real_name"`
	Password string    `json:"password"`
	OrgName  string    `json:"org_name"`
	OrgId    *int64    `json:"org_id"`
	LockFlag *int64    `json:"lock_flag"`
	Roles    []SysRole `gorm:"many2many:sys_user_role;"`
}

type SysRole struct {
	gorm.Model
	Name        string `json:"name" likeField:"name"`
	Code        string `json:"code"`
	Remarks     string `json:"remarks"`
	Description string `json:"description"`
	IsEdit      bool   `json:"-"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tpf_knowledge?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), gormConfig())
	if err != nil {
		fmt.Printf("%v", err)
	}
	user := make([]SysUser, 0)
	db.Preload("Roles").Limit(4).Offset(1).Find(&user)
	fmt.Printf("%v\n", user)
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
