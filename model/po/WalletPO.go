package po

import (
	"database/sql"
)

type WalletPO struct {
	Id         int          `db:"id"`
	Address    string       `db:"address"`
	Token      string       `db:"token"`
	LoginTime  sql.NullTime `db:"loginTime"`
	LogoutTime sql.NullTime `db:"logoutTime"`
	CreateTime sql.NullTime `db:"createTime"`
	UpdateTime sql.NullTime `db:"updateTime"`
}
