package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

var (
	mysqlConnection            = "mysql.connection"
	mysqlUserName              = "mysql.userName"
	mysqlPassword              = "mysql.password"
	mysqlPoolMaxConnection     = "mysql.pool.maxConnection"
	mysqlPoolCoreConnection    = "mysql.pool.coreConnection"
	mysqlConnectionIdleSeconds = "mysql.connection.idleSeconds"
	mysqlConnectionLifeSeconds = "mysql.connection.lifeSeconds"

	databaseType = "mysql"
)

var (
	MySqlClientBuildFaliure = errors.New("mysql client build failure")
)

func BuildMySqlClient() (*sqlx.DB, error) {
	connection := viper.GetString(mysqlConnection)
	userName := viper.GetString(mysqlUserName)
	password := viper.GetString(mysqlPassword)
	poolMaxConnection := viper.GetInt(mysqlPoolMaxConnection)
	poolCoreConnection := viper.GetInt(mysqlPoolCoreConnection)
	connectionIdleSeconds := viper.GetInt(mysqlConnectionIdleSeconds)
	connectionLifeSeconds := viper.GetInt(mysqlConnectionLifeSeconds)

	fullConnection := userName + ":" + password + "@" + connection

	db, err := sqlx.Open(databaseType, fullConnection)
	if err != nil {
		return nil, MySqlClientBuildFaliure
	}
	db.SetMaxOpenConns(poolMaxConnection)
	db.SetMaxIdleConns(poolCoreConnection)
	db.SetConnMaxLifetime(time.Duration(connectionLifeSeconds) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(connectionIdleSeconds) * time.Second)
	return db, nil
}
