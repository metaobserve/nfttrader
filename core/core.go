package core

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/dometa/component/cache"
	"github.com/dometa/component/database/mysql"
	"github.com/dometa/component/logger"
	"github.com/dometa/component/resources"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func BuildRuntimeCache() (*bigcache.BigCache, error) {

	cache, err := cache.BuildCache()
	if err != nil {
		fmt.Println("build runtime cache failure =>", err)
		return nil, err
	}
	fmt.Println("build runtime cache success")
	return cache, nil
}

// build environment first
func BuildEnvResources() error {
	err := resources.BuildEnvResources()
	if err != nil {
		return resources.ReadConfigurationFailure
	}
	fmt.Println("build env resources success")
	return nil
}

func BuildLogger() (*log.Logger, error) {
	log, err := logger.BuildLogger()
	if err != nil {
		fmt.Println("build logger failure =>", err)
		return nil, logger.LoggerBuildError
	}
	fmt.Println("build logger success")
	return log, nil
}

func BuildMySqlClient() (*sqlx.DB, error) {
	fmt.Println("build myql client")
	client, err := mysql.BuildMySqlClient()
	if err != nil {
		fmt.Println("build mysqlClient failure =>", err)
		return nil, err
	}
	fmt.Println("build mysqlClint success")
	return client, nil
}
