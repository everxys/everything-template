package vars

import (
	"path"
	"runtime"

	"everything-template/pkg/config"
	"github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	BasePath string
	Config   config.Config
	Redis    *redis.Client
	Lock     *redsync.Redsync
	DB       *gorm.DB
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(path.Dir(filename)))
	BasePath = root
}
