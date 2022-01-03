package components

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/universe-30/CliAppTemplate/cliCmd"
	"github.com/universe-30/Logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

/*
db_host
db_port
db_name
db_username
db_password
*/
func InitDB() (*gorm.DB, *sql.DB, error) {

	db_host, db_host_err := cliCmd.Config.GetString("db_host", "127.0.0.1")
	if db_host_err != nil {
		return nil, nil, errors.New("db_host [string] in config err," + db_host_err.Error())
	}

	db_port, db_port_err := cliCmd.Config.GetInt("db_port", 3306)
	if db_port_err != nil {
		return nil, nil, errors.New("db_port [int] in config err," + db_port_err.Error())
	}

	db_name, db_name_err := cliCmd.Config.GetString("db_name", "dbname")
	if db_name_err != nil {
		return nil, nil, errors.New("db_name [string] in config err," + db_name_err.Error())
	}

	db_username, db_username_err := cliCmd.Config.GetString("db_username", "username")
	if db_username_err != nil {
		return nil, nil, errors.New("db_username [string] in config err," + db_username_err.Error())
	}

	db_password, db_password_err := cliCmd.Config.GetString("db_password", "password")
	if db_password_err != nil {
		return nil, nil, errors.New("db_password [string] in config err," + db_password_err.Error())
	}

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&loc=UTC"

	GormDB, errOpen := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: New_gormLocalLogger(cliCmd.Logger),
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

///////////////////////////

type gormLocalLogger struct {
	LocalLogger           *Logrus.LocalLog
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func New_gormLocalLogger(localLogger *Logrus.LocalLog) *gormLocalLogger {
	return &gormLocalLogger{
		LocalLogger:           localLogger,
		SkipErrRecordNotFound: true,
	}
}

func (l *gormLocalLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *gormLocalLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	////when no err
	//if err == nil {
	//	//return when only interested in Error,Fatal,Panic
	//	if l.LocalLogger.Level < ULog_logrus.WarnLevel {
	//		return
	//	}
	//}
	//
	//elapsed := time.Since(begin)
	//if err == nil && l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
	//	//slow log
	//	sql, _ := fc()
	//	fields := ULog_logrus.Fields{}
	//	if l.SourceField != "" {
	//		fields[l.SourceField] = utils.FileWithLineNum()
	//	}
	//	l.LocalLogger.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
	//	return
	//}
	//
	/////errors , when error happens logs it at any loglevel
	//if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
	//	sql, _ := fc()
	//	fields := ULog_logrus.Fields{}
	//	if l.SourceField != "" {
	//		fields[l.SourceField] = utils.FileWithLineNum()
	//	}
	//	fields[ULog_logrus.ErrorKey] = err
	//	l.LocalLogger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
	//	return
	//}
	//
	////info
	//if l.LocalLogger.Level >= ULog_logrus.DebugLevel {
	//	sql, _ := fc()
	//	fields := ULog_logrus.Fields{}
	//	if l.SourceField != "" {
	//		fields[l.SourceField] = utils.FileWithLineNum()
	//	}
	//	if l.LocalLogger.Level == ULog_logrus.DebugLevel {
	//		l.LocalLogger.WithContext(ctx).WithFields(fields).Debugln("%s [%s]", sql, elapsed)
	//	} else {
	//		l.LocalLogger.WithContext(ctx).WithFields(fields).Traceln("%s [%s]", sql, elapsed)
	//	}
	//
	//}

}

func (l *gormLocalLogger) Info(ctx context.Context, s string, args ...interface{}) {
	//not used
}

func (l *gormLocalLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	//not used
}

func (l *gormLocalLogger) Error(ctx context.Context, s string, args ...interface{}) {
	//not used
}
