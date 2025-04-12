package config

import (
	"bytes"
	"fmt"

	"everything-template/configs"

	"github.com/spf13/viper"
)

func New(env string) *Config {
	v := viper.New()
	v.SetConfigType("yaml")

	data, err := configs.TemplateFs.ReadFile(fmt.Sprintf("config.%s.yaml", env))
	if err != nil {
		panic(fmt.Sprintf("read config failed", env, err))
	}

	err = v.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		panic(fmt.Sprintf("read config data error: %v", err))
	}

	cfg := &Config{}

	err = v.Unmarshal(cfg)
	if err != nil {
		panic(fmt.Sprintf("unmarshal config failed", env, err))
	}

	if len(env) > 0 {
		cfg.App.Env = env
	}

	return cfg
}
