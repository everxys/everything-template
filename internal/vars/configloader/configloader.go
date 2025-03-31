package configloader

import (
	"bytes"
	"everything-template/configs"
	"everything-template/internal/vars"
	"everything-template/internal/vars/config"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(env string) {
	v := viper.New()
	v.SetConfigType("yaml")

	data, err := configs.TemplateFs.ReadFile(fmt.Sprintf("config.%s.yaml", env))
	if err != nil {
		log.Fatalf("read config %s: %v", env, err)
	}

	err = v.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("read config data error: %v", err)
	}

	cfg := &config.Config{}

	err = v.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("Unmarshal config data error: %v", err)
	}

	if len(env) > 0 {
		cfg.App.Env = env
	}

	vars.Config = cfg
}
