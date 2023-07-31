package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm" //导入gorm框架的包
	"gorm.io/gorm/logger"
	"log"
	"questionnaire/config"
)

var DB *gorm.DB //定义全局变量，用于存储数据库连接

func InitDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DB.UserName, //从配置文件中获取数据库登录用户的信息
		config.Config.DB.Password, //从配置文件中获取数据库登录密码信息
		config.Config.DB.Address,  //从配置文件中获取数据库地址信息
		config.Config.DB.DBName,   //从配置文件中获取数据库名称信息
	)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ //连接数据库
		DisableForeignKeyConstraintWhenMigrating: true,                                // 关闭外键约束 提升数据库速度
		Logger:                                   logger.Default.LogMode(logger.Info), //指定日志级别为Info级别
	})
	if err != nil {
		log.Fatal("DatabaseConnectFailed", err)
	}
	err = autoMigrate(db)
	if err != nil {
		log.Fatal("DatabaseMigrateFailed", err)
	}
	DB = db
}
