package logs

import (
	"github.com/universe-30/UCliAppTemplate/cli"
)

func StartLog() {
	num := cli.CmdToDo.CliContext.Int("num")
	if num == 0 {
		num = 20
	}

	onlyerr := cli.CmdToDo.CliContext.Bool("onlyerr")
	if onlyerr {
		cli.Logger.PrintLastN_ErrLogs(num)
	} else {
		cli.Logger.PrintLastN_AllLogs(num)
	}
}
