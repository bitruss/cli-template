package default_

import (
	"github.com/coreservice-io/CliAppTemplate/cmd/default_/http"
)

func start_service() {

	//start the httpserver
	http.StartHttpSever()
}
