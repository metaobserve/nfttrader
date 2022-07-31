package cache

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	time "time"
)

var (
	CacheBuildError        = errors.New("runtime cache was built in error.")
	runtimeExpireInMinutes = "cache.runtime.expireInMinutes"
)

func BuildCache() (*bigcache.BigCache, error) {
	fmt.Println("build runtime cache")

	expireTime := viper.GetInt(runtimeExpireInMinutes)
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(expireTime) * time.Minute))
	if err != nil {
		fmt.Println("runtime cache build throw err => ", err)
		return nil, CacheBuildError
	}
	return cache, nil
}
