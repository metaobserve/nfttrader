package po

import "database/sql"

type WhiteListPO struct {
	Id             int          `db:"id"`
	Name           string       `db:"name"`
	Description    string       `db:"description"`
	StartTime      sql.NullTime `db:"startTime"`
	EndTime        sql.NullTime `db:"endTime"`
	AirdropTokenId string       `db:"airdropTokenId"`
	// status : 0 invalid , 1 valid
	Status        int          `db:"status"`
	WalletAddress string       `db:"walletAddress"`
	CreateTime    sql.NullTime `db:"createTime"`
	UpdateTime    sql.NullTime `db:"updateTime"`
}
