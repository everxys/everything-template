package boots

import (
	"everything-template/internal/vars"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	vars.Redis = redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%s", vars.Config.Redis.Addr, vars.Config.Redis.Port),
		Password:        vars.Config.Redis.Password,
		DB:              vars.Config.Redis.DB,
		PoolSize:        vars.Config.Redis.PoolSize,
		MinIdleConns:    vars.Config.Redis.MinIdleConn,
		MaxIdleConns:    vars.Config.Redis.MaxIdleConn,
		ConnMaxLifetime: vars.Config.Redis.MaxLifeTime * time.Second,
		ConnMaxIdleTime: vars.Config.Redis.MaxIdleTime * time.Minute,
	})
}
