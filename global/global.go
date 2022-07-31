package global

import (
	"github.com/allegro/bigcache/v3"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	Logger      *log.Logger
	MysqlClient *sqlx.DB
	Cache       *bigcache.BigCache
)
