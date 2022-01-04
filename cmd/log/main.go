package log

import (
	"github.com/universe-30/CliAppTemplate/cliCmd"
	"github.com/universe-30/ULog"
)

func StartLog() {
	num := cliCmd.CmdToDo.CliContext.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := cliCmd.CmdToDo.CliContext.Bool("onlyerr")
	if onlyerr {
		cliCmd.Logger.PrintLastN(num, []ULog.LogLevel{ULog.PanicLevel, ULog.FatalLevel, ULog.ErrorLevel})
	} else {
		cliCmd.Logger.PrintLastN(num, []ULog.LogLevel{ULog.PanicLevel, ULog.FatalLevel, ULog.ErrorLevel, ULog.InfoLevel, ULog.WarnLevel, ULog.DebugLevel, ULog.TraceLevel})
	}
}
