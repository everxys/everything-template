package bootstrap

import (
	"everything-template/internal/bootstrap/boots"
	"everything-template/internal/server"
	"everything-template/internal/vars/configloader"
)

func Run(env string) {
	configloader.LoadConfig(env)

	boots.InitPostgres()
	boots.InitRedis()

	server.Run()
}
