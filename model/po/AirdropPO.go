package po

import (
	"database/sql"
)

type AirdropPO struct {
	Id int `db:"id"`
	// airdrop token, guid value
	TokenId     string       `db:"tokenId"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	StartTime   sql.NullTime `db:"startTime"`
	EndTime     sql.NullTime `db:"endTime"`
	// airDrop stage: 0 ready; 1 action; 2 cut
	Stage      int          `db:"stage"`
	CreateTime sql.NullTime `db:"createTime"`
	UpdateTime sql.NullTime `db:"updateTime"`
}
