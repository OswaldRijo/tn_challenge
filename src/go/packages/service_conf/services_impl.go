package service_conf

import (
	"truenorth/packages/utils"
)

var (
	localServer *utils.ServerStartUp
)

func initLocalServerInstance() {
	localServer = new(utils.ServerStartUp)
}
func GetLocalServerInstance() *utils.ServerStartUp {
	if localServer == nil {
		initLocalServerInstance()
	}

	return localServer
}
