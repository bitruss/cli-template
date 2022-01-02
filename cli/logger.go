package cli

import (
	"github.com/daqnext/utils/path_util"
	"github.com/fatih/color"
	"github.com/universe-30/ULog"
	"github.com/universe-30/ULog_logrus"
)

var Logger *ULog_logrus.LocalLog

func InitLogger() {
	var llerr error
	Logger, llerr = ULog_logrus.New(path_util.GetAbsPath("logs"), 2, 20, 30)

	if llerr != nil {
		color.Set(color.FgRed)
		defer color.Unset()
		panic("Error:" + llerr.Error())
	}
}

func SetLogLevel(logLevel string) {
	l := ULog.LogLevelStrToLevel(logLevel)
	Logger.SetLevel(l)

	if l == ULog.DebugLevel || l == ULog.TraceLevel {
		Logger.SetReportCaller(true)
	}
}
