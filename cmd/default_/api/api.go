package api

import "github.com/coreservice-io/CliAppTemplate/plugin/echoServer"

func DeclareApi(httpServer *echoServer.EchoServer) {
	//health
	httpServer.GET("/api/health", healthHandler)
}
