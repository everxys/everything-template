package boots

import (
	"everything-template/internal/vars"
	"everything-template/pkg/config"
)

func InitConfig(env string) {
	cfg := config.New(env)
	vars.Config = *cfg
}
