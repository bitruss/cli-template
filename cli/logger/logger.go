package logger

import (
	"github.com/daqnext/utils/path_util"
	"github.com/fatih/color"
	"github.com/universe-30/ULog"
)

var LocalLogger *ULog.LocalLog

func IniLocalLogger() {
	var llerr error
	LocalLogger, llerr = ULog.New(path_util.GetAbsPath("logs"), 2, 20, 30)

	if llerr != nil {
		color.Set(color.FgRed)
		defer color.Unset()
		panic("Error:" + llerr.Error())
	}
}

func SetLogLevel(logLevel ULog.LogLevel) error {
	err := LocalLogger.ResetLevel(logLevel)
	if err != nil {
		return err
	}

	if logLevel == ULog.LEVEL_DEBUG || logLevel == ULog.LEVEL_TRACE {
		LocalLogger.SetReportCaller(true)
	}
	return nil
}
