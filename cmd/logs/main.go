package logs

import (
	"github.com/universe-30/CliAppTemplate/cliCmd"
)

func StartLog() {
	num := cliCmd.CmdToDo.CliContext.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := cliCmd.CmdToDo.CliContext.Bool("onlyerr")
	if onlyerr {
		cliCmd.Logger.PrintLastN_ErrLogs(num)
	} else {
		cliCmd.Logger.PrintLastN_AllLogs(num)
	}
}
