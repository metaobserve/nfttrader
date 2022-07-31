package po

import (
	"database/sql"
)

type MintPO struct {
	Id         int          `db:"id"`
	Address    string       `db:"address"`
	NftId      int          `db:"nftId"`
	AirdropId  int          `db:"airdropId""`
	CreateTime sql.NullTime `db:"createTime"`
	UpdateTime sql.NullTime `db:"updateTime"`
}
