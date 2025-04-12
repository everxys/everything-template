package boots

import (
	"context"
	"fmt"
	"time"

	"everything-template/internal/vars"
	"everything-template/pkg/util"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:            util.NewFluentBuilder().WriteString(vars.Config.Redis.Addr).WriteString(":").WriteString(vars.Config.Redis.Port).String(),
		Password:        vars.Config.Redis.Password,
		DB:              vars.Config.Redis.DB,
		PoolSize:        vars.Config.Redis.PoolSize,
		MinIdleConns:    vars.Config.Redis.MinIdleConn,
		MaxIdleConns:    vars.Config.Redis.MaxIdleConn,
		ConnMaxLifetime: vars.Config.Redis.MaxLifeTime * time.Second,
		ConnMaxIdleTime: vars.Config.Redis.MaxIdleTime * time.Minute,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Sprintf("Redis connection failed: %v", err))
	}

	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	vars.Redis = client
	vars.Lock = rs
}
