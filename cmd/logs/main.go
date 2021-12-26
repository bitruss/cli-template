package logs

import (
	"github.com/universe-30/UCliAppTemplate/cli"
	"github.com/universe-30/UCliAppTemplate/cli/logger"
)

func StartLog() {
	num := cli.CmdToDo.CliContext.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := cli.CmdToDo.CliContext.Bool("onlyerr")
	if onlyerr {
		logger.LocalLogger.PrintLastN_ErrLogs(num)
	} else {
		logger.LocalLogger.PrintLastN_AllLogs(num)
	}
}
