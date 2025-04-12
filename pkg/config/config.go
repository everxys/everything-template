package config

import (
	"time"
)

type Config struct {
	App      AppConfig      `yaml:"app" json:"app"`
	Redis    RedisConfig    `yaml:"redis" json:"redis"`
	Postgres PostgresConfig `yaml:"postgres" json:"postgres"`
	Auth     AuthConfig     `yaml:"auth" json:"auth"`
}

type AppConfig struct {
	Name string `yaml:"name" json:"name"`
	Port int    `yaml:"port" json:"port"`
	Env  string `yaml:"env" json:"env"`
}

type RedisConfig struct {
	Addr        string        `yaml:"addr" json:"addr"`
	Port        string        `yaml:"port" json:"port"`
	Password    string        `yaml:"password" json:"password"`
	DB          int           `yaml:"db" json:"db"`
	PoolSize    int           `yaml:"poolSize" json:"poolSize"`
	MinIdleConn int           `yaml:"minIdleConn" json:"minIdleConn"`
	MaxIdleConn int           `yaml:"maxIdleConn" json:"maxIdleConn"`
	MaxLifeTime time.Duration `yaml:"maxLifeTime" json:"maxLifeTime"`
	MaxIdleTime time.Duration `yaml:"maxIdleTime" json:"maxIdleTime"`
}

type PostgresConfig struct {
	Addr     string `yaml:"addr" json:"addr"`
	Port     int    `yaml:"port" json:"port"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Dbname   string `yaml:"dbname" json:"dbname"`
}

type AuthConfig struct {
	TokenExpiry  time.Duration `yaml:"tokenExpiry" json:"tokenExpiry"`
	CookieMaxAge int           `yaml:"cookieMaxAge" json:"cookieMaxAge"`
	SecretKey    string        `yaml:"secretKey" json:"secretKey"`
}
