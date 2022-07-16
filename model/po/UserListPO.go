package po

import (
	"database/sql"
)

type UserListPO struct {
	Id              int          `db:"id"`
	WalletId        int          `db:"walletId"`
	AirdropId       int          `db:"airdropId"`
	JoinedWhiteList bool         `db:"joinedWhiteList"`
	CreateTime      sql.NullTime `db:"createTime"`
	UpdateTime      sql.NullTime `db:"updateTime"`
}
