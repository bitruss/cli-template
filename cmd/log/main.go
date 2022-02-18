package log

import (
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/ULog"
	"github.com/urfave/cli/v2"
)

func StartLog(clictx *cli.Context) {
	num := clictx.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := clictx.Bool("only_err")
	if onlyerr {
		basic.Logger.PrintLastN(num, []ULog.LogLevel{ULog.PanicLevel, ULog.FatalLevel, ULog.ErrorLevel})
	} else {
		basic.Logger.PrintLastN(num, []ULog.LogLevel{ULog.PanicLevel, ULog.FatalLevel, ULog.ErrorLevel, ULog.InfoLevel, ULog.WarnLevel, ULog.DebugLevel, ULog.TraceLevel})
	}
}
