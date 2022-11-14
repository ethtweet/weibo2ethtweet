package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"weibo2ipfs/model"
)

var Db = &dbUtils{
	tables: []interface{}{
		&model.User{},
		&model.Weibo{},
	},
}

type dbUtils struct {
	tables []interface{}
	db     *gorm.DB
}

// AutoMigrate 自动迁移
func (s *dbUtils) AutoMigrate() error {
	return s.Get().AutoMigrate(s.tables...)
}

func (s *dbUtils) Get() *gorm.DB {
	if s.db != nil {
		return s.db
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_DATABASE"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,        // 禁用彩色打印
		},
	)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		// 禁用 AutoMigrate 自动创建数据库外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	s.db = _db
	return s.db
}
