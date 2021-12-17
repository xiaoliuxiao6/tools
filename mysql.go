package tools

import (
	"encoding/json"
	"fmt"
	logos "log"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ----------------------------------------------------------------------------- 初始化数据库连接
func InitDB() *gorm.DB {
	// 连接字符串定义
	dbHost := "10.100.0.29"
	dbPort := "3306"
	dbUser := "wax"
	dbPassword := "Xiaoliu123.com&*&*"
	dbName := "opensea"

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=UTC", dbUser, dbPassword, dbHost, dbPort, dbName)

	newLogger := logger.New(
		logos.New(os.Stdout, "\r\n", logos.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 彩色打印
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // 打印所有 SQL 语句
		Logger: newLogger,
	})

	// 错误处理
	if err != nil {
		log.Panicln("数据库连接失败", dsn)
	}

	// 自动迁移
	// db.AutoMigrate(&Assets{})

	// 连接池设置
	// 使用连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicln("设置连接池失败", dsn)
	}

	sqlDB.SetMaxIdleConns(2)            //设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(10)           //设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour) //设置了连接可复用的最大时间。

	// 确保服务可用
	// Ping
	err = sqlDB.Ping()
	if err != nil {
		log.Panicln("Ping Error")
	}

	// 获取数据库配置
	data, _ := json.Marshal(sqlDB.Stats())
	log.Info(string(data))

	// 打印数据库信息
	// status := sqlDB.Stats()
	// log.Infof("最大连接数：%v, 打开连接数：%v", status.MaxOpenConnections, status.OpenConnections)
	return db
}

// ----------------------------------------------------------------------------- 定义全局数据库变量
// var db = initDB()
