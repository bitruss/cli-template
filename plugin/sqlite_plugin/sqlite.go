package sqlite_plugin

import (
	"fmt"
	"time"

	"github.com/coreservice-io/GormULog"
	"github.com/coreservice-io/ULog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instanceMap = map[string]*gorm.DB{}

func GetInstance() *gorm.DB {
	return instanceMap["default"]
}

func GetInstance_(name string) *gorm.DB {
	return instanceMap[name]
}

type Config struct {
	Sqlite_path string
}

func Init(dbConfig Config, logger ULog.Logger) error {
	return Init_("default", dbConfig, logger)
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init_(name string, dbConfig Config, logger ULog.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("sqlite instance <%s> has already been initialized", name)
	}

	//Level: Silent Error Warn Info. Info logs all record. Silent turns off log.
	db_log_level := GormULog.Warn
	if logger.GetLevel() >= ULog.TraceLevel {
		db_log_level = GormULog.Info
	}

	db, err := gorm.Open(sqlite.Open(dbConfig.Sqlite_path), &gorm.Config{
		Logger: GormULog.New_gormLocalLogger(logger, GormULog.Config{
			SlowThreshold:             500 * time.Millisecond,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  db_log_level,
		}),
	})

	if err != nil {
		return err
	}

	_, err = db.DB()
	if err != nil {
		return err
	}
	//sqlDB.SetMaxIdleConns(5)
	//sqlDB.SetMaxOpenConns(20)

	instanceMap[name] = db

	return nil
}
