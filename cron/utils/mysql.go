package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MysqlDefaultDb *gorm.DB

// MysqlLogDb 日志库连接语柄
var MysqlLogDb *gorm.DB

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func InitMysql() (err error) {
	MysqlDefaultDb, err = gorm.Open(mysql.Open(getDsn("platform")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "lhc_", // 表前缀
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		Logger:                   logger.Default.LogMode(logger.Warn), // 日志等级
		DisableNestedTransaction: true,                                // 禁止自动创建外键
	})
	if err != nil {
		return err
	}

	MysqlLogDb, err = gorm.Open(mysql.Open(getDsn("log")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "lhc_", // 表前缀
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		Logger:                   logger.Default.LogMode(logger.Warn), // 日志等级
		DisableNestedTransaction: true,                                // 禁止自动创建外键
		// SkipDefaultTransaction:   true,                                // 关闭事务
	})
	if err != nil {
		return err
	}

	return nil
}

func getDsn(db string) string {
	var config = MysqlConfig{
		Host:     viper.GetString("mysql." + db + ".master.host"),
		Port:     viper.GetInt("mysql." + db + ".master.port"),
		User:     viper.GetString("mysql." + db + ".master.user"),
		Password: viper.GetString("mysql." + db + ".master.password"),
		Database: viper.GetString("mysql." + db + ".master.database"),
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database,
	)
}
