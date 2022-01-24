package sqldb

import (
	"database/sql"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/configuration"
	"github.com/universe-30/GormULog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func Init() {
	//only run once
	once.Do(func() {
		var err error = nil
		db, _, err = newDB()
		if err != nil {
			basic.Logger.Fatalln(err)
		}
	})
}

func GetSingleInstance() *gorm.DB {
	Init()
	return db
}

/*
db_host
db_port
db_name
db_username
db_password
*/
func newDB() (*gorm.DB, *sql.DB, error) {

	db_host, db_host_err := configuration.Config.GetString("db_host", "127.0.0.1")
	if db_host_err != nil {
		return nil, nil, errors.New("db_host [string] in config err," + db_host_err.Error())
	}

	db_port, db_port_err := configuration.Config.GetInt("db_port", 3306)
	if db_port_err != nil {
		return nil, nil, errors.New("db_port [int] in config err," + db_port_err.Error())
	}

	db_name, db_name_err := configuration.Config.GetString("db_name", "dbname")
	if db_name_err != nil {
		return nil, nil, errors.New("db_name [string] in config err," + db_name_err.Error())
	}

	db_username, db_username_err := configuration.Config.GetString("db_username", "username")
	if db_username_err != nil {
		return nil, nil, errors.New("db_username [string] in config err," + db_username_err.Error())
	}

	db_password, db_password_err := configuration.Config.GetString("db_password", "password")
	if db_password_err != nil {
		return nil, nil, errors.New("db_password [string] in config err," + db_password_err.Error())
	}

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&loc=UTC"

	var GormDB *gorm.DB
	var errOpen error

	GormDB, errOpen = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: GormULog.New_gormLocalLogger(basic.Logger, GormULog.Config{
			SlowThreshold:             500 * time.Millisecond,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  GormULog.Warn, //Level: Silent Error Warn Info. Info logs all record. Silent turns off log.
		}),
	})

	if errOpen != nil {
		return nil, nil, errOpen
	}

	sqlDB, errsql := GormDB.DB()
	if errsql != nil {
		return nil, nil, errsql
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)

	return GormDB, sqlDB, nil

}
