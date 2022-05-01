package basic

import (
	"github.com/coreservice-io/LogrusULog"
	"github.com/coreservice-io/ULog"
	"github.com/coreservice-io/UUtils/path_util"
	"github.com/fatih/color"
)

var Logger ULog.Logger

func InitLogger() {
	var llerr error
	logs_path := path_util.ExE_Path("logs")
	Logger, llerr = LogrusULog.New(logs_path, 2, 20, 30)

	if llerr != nil {
		color.Set(color.FgRed)
		defer color.Unset()
		panic("Error:" + llerr.Error())
	}

	Logger.Infoln("logs_path:", logs_path)
}
