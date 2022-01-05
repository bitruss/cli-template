package log

import (
	"github.com/universe-30/CliAppTemplate/boot"
	"github.com/universe-30/ULog"
)

func StartLog() {
	num := boot.CmdToDo.CliContext.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := boot.CmdToDo.CliContext.Bool("onlyerr")
	if onlyerr {
		boot.Logger.PrintLastN(num, []ULog.LogLevel{ULog.PanicLevel, ULog.FatalLevel, ULog.ErrorLevel})
	} else {
		boot.Logger.PrintLastN(num, []ULog.LogLevel{ULog.PanicLevel, ULog.FatalLevel, ULog.ErrorLevel, ULog.InfoLevel, ULog.WarnLevel, ULog.DebugLevel, ULog.TraceLevel})
	}
}
