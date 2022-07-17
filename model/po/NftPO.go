package po

import (
	"database/sql"
)

type NftPO struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	// nft token
	Token          string `db:"token"`
	AirdropTokenId string `db:"airdropTokenId"`
	// nft status : 0 invalid 1 valid
	Status     int          `db:"status"`
	CreateTime sql.NullTime `db:"createTime"`
	UpdateTime sql.NullTime `db:"updateTime"`
}
