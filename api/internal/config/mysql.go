package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MysqlDefaultDb *gorm.DB

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func InitMysql() (err error) {
	MysqlDefaultDb, err = gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "lhc_", // 表前缀
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		Logger:                   logger.Default.LogMode(logger.Info), // 日志等级
		DisableNestedTransaction: true,                                // 禁止自动创建外键
	})
	if err != nil {
		return err
	}

	return nil
}

func getDsn() string {
	var config = MysqlConfig{
		Host:     viper.GetString("mysql.master.host"),
		Port:     viper.GetInt("mysql.master.port"),
		User:     viper.GetString("mysql.master.user"),
		Password: viper.GetString("mysql.master.password"),
		Database: viper.GetString("mysql.master.database"),
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database,
	)
}
