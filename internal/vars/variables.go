package vars

import (
	"everything-template/pkg/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"path"
	"runtime"
)

var (
	BasePath string
	Config   config.Config
	Redis    *redis.Client
	DB       *gorm.DB
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(path.Dir(filename)))
	BasePath = root
}
