package boot

import (
	"os"
)

func ReadArgs() {
	//config app to run
	errRun := configCmd().Run(os.Args)
	if errRun != nil {
		Logger.Panicln(errRun)
	}
}
