package initialize

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/pkg/tracer"
)

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		global.ServerConfig.MysqlInfo.Username, global.ServerConfig.MysqlInfo.Password,
		global.ServerConfig.MysqlInfo.Host, global.ServerConfig.MysqlInfo.Port,
		global.ServerConfig.MysqlInfo.Database)

	var ormLogger logger.Interface
	if global.ServerConfig.LogLevel == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: ormLogger,
	})
	if err != nil {
		return err
	}

	_ = db.Use(&tracer.OpentracingPlugin{})
	global.DB = db
	return nil
}

/*
func InitDB() error {
	var err error

	//question: unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	//answer: https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	global.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		global.ServerConfig.MysqlInfo.Username, global.ServerConfig.MysqlInfo.Password,
		global.ServerConfig.MysqlInfo.Host, global.ServerConfig.MysqlInfo.Port,
		global.ServerConfig.MysqlInfo.Database))
	if err != nil {
		return err
	}

	global.DB.SetMaxOpenConns(100) // 设置最大连接数
	global.DB.SetMaxIdleConns(100) // 设置最大空闲连接数
	return nil
}
*/
