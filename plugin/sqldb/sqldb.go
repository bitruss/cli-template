package sqldb

import (
	"errors"
	"strconv"
	"time"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/configuration"
	"github.com/universe-30/GormULog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetSingleInstance() *gorm.DB {
	return db
}

/*
db_host
db_port
db_name
db_username
db_password
*/
func Init() error {
	if db != nil {
		return nil
	}
	db_host, err := configuration.Config.GetString("db_host", "127.0.0.1")
	if err != nil {
		return errors.New("db_host [string] in config err," + err.Error())
	}

	db_port, err := configuration.Config.GetInt("db_port", 3306)
	if err != nil {
		return errors.New("db_port [int] in config err," + err.Error())
	}

	db_name, err := configuration.Config.GetString("db_name", "dbname")
	if err != nil {
		return errors.New("db_name [string] in config err," + err.Error())
	}

	db_username, err := configuration.Config.GetString("db_username", "username")
	if err != nil {
		return errors.New("db_username [string] in config err," + err.Error())
	}

	db_password, err := configuration.Config.GetString("db_password", "password")
	if err != nil {
		return errors.New("db_password [string] in config err," + err.Error())
	}

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&loc=UTC"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: GormULog.New_gormLocalLogger(basic.Logger, GormULog.Config{
			SlowThreshold:             500 * time.Millisecond,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  GormULog.Warn, //Level: Silent Error Warn Info. Info logs all record. Silent turns off log.
		}),
	})

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)

	return nil
}
