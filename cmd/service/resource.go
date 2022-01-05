package service

import "github.com/universe-30/CliAppTemplate/components"

//define your components here

var CompDeamon *components.Service

func iniResources() {
	CompDeamon = components.NewDaemonService()

}

func releaseResources() {
	//DaemonService.Stop()
}
