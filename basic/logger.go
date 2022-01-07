package basic

import (
	"github.com/fatih/color"
	"github.com/universe-30/LogrusULog"
	"github.com/universe-30/ULog"
	"github.com/universe-30/UUtils/path_util"
)

var Logger ULog.Logger

func InitLogger() {
	var llerr error
	Logger, llerr = LogrusULog.New(path_util.GetAbsPath("logs"), 2, 20, 30)

	if llerr != nil {
		color.Set(color.FgRed)
		defer color.Unset()
		panic("Error:" + llerr.Error())
	}
}

func SetLogLevel(logLevel string) {
	l := ULog.ParseLogLevel(logLevel)
	Logger.SetLevel(l)
}
