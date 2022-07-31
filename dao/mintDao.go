package dao

import (
	"database/sql"
	"github.com/dometa/global"
	"github.com/dometa/model/po"
)

type MintDao interface {
	Insert(mint po.MintPO) (bool, error)
	Get(mint *po.MintPO, address string) error
}

type mintDaoImpl struct {
}

func NewMintDao() MintDao {
	return &mintDaoImpl{}
}

func (dao mintDaoImpl) Insert(mint po.MintPO) (bool, error) {
	sqlStr := "insert into mint(address,nftId,airdropId,createTime,updateTime)values(?,?,?,?,?)"
	ret, err := global.MysqlClient.Exec(sqlStr,
		mint.Address,
		mint.NftId,
		mint.AirdropId,
		mint.CreateTime,
		mint.UpdateTime,
	)
	if err != nil {
		global.Logger.
			WithField("dao", "insertMintError").Errorln("insert mint error", err)
		return false, InsertError
	}
	affectRows, err := ret.RowsAffected()
	if err != nil {
		global.Logger.WithField("add mint", "noAffectedError").WithField("mint", mint).Errorln("no affect error=>", err)
		return false, InsertError
	}
	if affectRows > 0 {
		return true, nil
	} else {
		return false, NothingInsertError
	}
}

func (dao mintDaoImpl) Get(mint *po.MintPO, address string) error {
	err := global.MysqlClient.Get(mint, "select * from mint where address = ?", address)
	if err != nil {
		global.Logger.
			WithField("dao", "getMintError").Errorln("get mint error", err)
		if err == sql.ErrNoRows {
			return nil
		}
		return SelectError
	} else {
		return nil
	}
}
