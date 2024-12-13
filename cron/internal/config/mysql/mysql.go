package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var PlatformDB *gorm.DB

// LogDb  日志库连接语柄
var LogDb *gorm.DB

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func Init() (err error) {
	PlatformDB, err = gorm.Open(mysql.Open(getDsn("platform")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "lhc_", // 表前缀
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		Logger:                   logger.Default.LogMode(logger.Error), // 日志等级
		DisableNestedTransaction: true,                                 // 禁止自动创建外键
	})
	if err != nil {
		return err
	}

	LogDb, err = gorm.Open(mysql.Open(getDsn("log")), &gorm.Config{
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
	var config = Config{
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
