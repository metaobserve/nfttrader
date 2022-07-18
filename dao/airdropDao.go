package dao

import (
	"github.com/dometa/global"
	"github.com/dometa/model/po"
	"github.com/jmoiron/sqlx"
)

type airdropDao interface {
	// get the airdrop by stage. There is only one action airport everytime.
	GetAirdropByStage(airdrop *po.AirdropPO, stage int) error
	Insert(airport po.AirdropPO) (bool, error)
}

type airdropDaoImpl struct {
	sqlClient *sqlx.DB
}

func NewAirportDao(sqlClient *sqlx.DB) airdropDao {
	return &airdropDaoImpl{
		sqlClient: sqlClient,
	}
}

func (dao airdropDaoImpl) Insert(airdrop po.AirdropPO) (bool, error) {
	sqlStr := "insert into airdrop(tokenId,name,description,startTime,endTime,stage,createTime,updateTime)values(?,?,?,?,?,?,?,?)"
	ret, err := dao.sqlClient.Exec(sqlStr, airdrop.TokenId, airdrop.Name, airdrop.Description, airdrop.StartTime, airdrop.EndTime, airdrop.Stage, airdrop.CreateTime, airdrop.UpdateTime)
	if err != nil {
		global.Logger.
			WithField("insertAirdrop", "error").
			WithField("airdrop", airdrop).Errorln("insert airdrop error", err)
		return false, InsertError
	}
	affectRows, err := ret.RowsAffected()
	if err != nil {
		global.Logger.WithField("insertAirdrop", "noAffectedError").WithField("airdrop", airdrop).Errorln("no affect error")
		return false, InsertError
	}
	if affectRows > 0 {
		return true, nil
	} else {
		return false, NothingInsertError
	}
}

func (dao airdropDaoImpl) GetAirdropByStage(airdrop *po.AirdropPO, stage int) error {
	airdrop.StartTime.Valid = true
	airdrop.UpdateTime.Valid = true
	airdrop.CreateTime.Valid = true
	airdrop.EndTime.Valid = true
	err := dao.sqlClient.Get(airdrop, "select * from airdrop where stage = ?", stage)
	if err != nil {
		global.Logger.WithField("getAirdrop", "error").WithField("stage", stage).Errorln("get airdrop by state error")
		return SelectError
	}
	return nil
}
