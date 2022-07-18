package global

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	Logger      *log.Logger
	MysqlClient *sqlx.DB
)