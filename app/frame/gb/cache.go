package gb

import (
	"github.com/donetkit/aqara.sdk/app/config"
	"github.com/donetkit/aqara.sdk/app/frame/cache"
)

var Cache *cache.Cacher

func RegisterCache() *cache.Cacher {
	c, err := cache.New(
		cache.Options{
			Prefix:   "",
			Addr:     config.RedisConfig.Addr,
			Port:     config.RedisConfig.Port,
			Password: config.RedisConfig.Password,
			Db:       config.RedisConfig.DB,
		})
	if err != nil {
		panic(err)
	}
	Cache = c

	return c
}
