package basic

import (
	"github.com/coreservice-io/log"
	"github.com/coreservice-io/logrus_log"
	"github.com/fatih/color"
)

var Logger log.Logger

func InitLogger() {
	var llerr error

	logs_abs_path := AbsPath("logs")
	Logger, llerr = logrus_log.New(logs_abs_path, 2, 20, 30)

	if llerr != nil {
		color.Set(color.FgRed)
		defer color.Unset()
		panic("Error:" + llerr.Error())
	}

	Logger.Infoln("logs_path:", logs_abs_path)
}
