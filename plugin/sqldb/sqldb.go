package sqldb

import (
	"fmt"
	"strconv"
	"time"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/GormULog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instanceMap = map[string]*gorm.DB{}

func GetDefaultInstance() *gorm.DB {
	return instanceMap["default"]
}

func GetInstance(name string) *gorm.DB {
	return instanceMap[name]
}

/*
db_host
db_port
db_name
db_username
db_password
*/
type Config struct {
	Host     string
	Port     int
	DbName   string
	UserName string
	Password string
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init(name string, dbConfig Config) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("db instance <%s> has already initialized", name)
	}

	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbConfig.DbName + "?charset=utf8mb4&loc=UTC"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

	instanceMap[name] = db

	return nil
}
