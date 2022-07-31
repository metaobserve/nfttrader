package po

import (
	"database/sql"
)

type NftPO struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	// nft token
	//Token          string `db:"token"`
	//AirdropTokenId int `db:"airdropTokenId"`
	// nft status : 0 invalid 1 valid
	Status        int          `db:"status"`
	Theme         string       `db:"theme"`
	Category      string       `db:"category"`
	Price         string       `db:"price"`
	PriceUnit     string       `db:"priceUnit"`
	NftAddress    string       `db:"nftAddress"`
	Author        string       `db:"author"`
	AuthorAddress string       `db:"authorAddress"`
	CreateTime    sql.NullTime `db:"createTime"`
	UpdateTime    sql.NullTime `db:"updateTime"`
}
