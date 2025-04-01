package bootstrap

import (
	"everything-template/internal/bootstrap/boots"
	"everything-template/internal/server"
)

func Run(env string) {

	boots.InitConfig(env)
	boots.InitPostgres()
	boots.InitRedis()

	server.Run()
}
