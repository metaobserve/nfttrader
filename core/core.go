package core

import (
	"fmt"
	"github.com/dometa/component/database/mysql"
	"github.com/dometa/component/logger"
	"github.com/dometa/component/resources"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// build environment first
func BuildEnvResources() error {
	err := resources.BuildEnvResources()
	if err != nil {
		return resources.ReadConfigurationFailure
	}
	return nil
}

func BuildLogger() (*log.Logger, error) {
	fmt.Println("build a logger and inject into all the layers.")
	log, err := logger.BuildLogger()
	if err != nil {
		fmt.Println("build logger failure =>", err)
		return nil, logger.LoggerBuildError
	}
	log.
		WithField("build", "logger").
		Infoln("build logger success")
	return log, nil
}

func BuildMySqlClient() (*sqlx.DB, error) {
	fmt.Println("build myql client")
	client, err := mysql.BuildMySqlClient()
	if err != nil {
		fmt.Println("build mysqlClient failure =>", err)
		return nil, err
	}
	return client, nil
}
