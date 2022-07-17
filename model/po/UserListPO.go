package po

import (
	"database/sql"
)

type UserListPO struct {
	Id              int          `db:"id"`
	WalletAddress   string       `db:"walletAddress"`
	AirdropTokenId  string       `db:"airdropTokenId"`
	JoinedWhiteList bool         `db:"joinedWhiteList"`
	CreateTime      sql.NullTime `db:"createTime"`
	UpdateTime      sql.NullTime `db:"updateTime"`
}
