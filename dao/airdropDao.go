package dao

import (
	"github.com/dometa/bootstrap"
	"github.com/dometa/model/po"
)

type airdropDao interface {
	// get the airdrop by stage. There is only one action airport everytime.
	GetAirdropByStage(airdrop *po.AirdropPO, stage int) error
	Insert(airport po.AirdropPO) (bool, error)
}

type airdropDaoImpl struct {
	context *bootstrap.Context
}

func NewAirportDao(context *bootstrap.Context) airdropDao {
	return &airdropDaoImpl{
		context: context,
	}
}

func (dao airdropDaoImpl) Insert(airport po.AirdropPO) (bool, error) {
	sqlStr := "insert into airport(name,description,startTime,endTime,stage,createTime,updateTime)values(?,?,?,?,?,?,?)"
	result, err := dao.context.MysqlClient.Exec(sqlStr, airport.Name, airport.Description, airport.StartTime, airport.EndTime, airport.CreateTime, airport.UpdateTime)
	if err != nil {
		return false, FailInsertError
	}
	affectedNum, err := result.RowsAffected()
	if err != nil {
		return false, FailInsertError
	}
	if affectedNum > 0 {
		return true, nil
	} else {
		return false, NoneInsertError
	}
}

func (dao airdropDaoImpl) GetAirdropByStage(airdrop *po.AirdropPO, stage int) error {
	err := dao.context.MysqlClient.Get(airdrop, "select * from airport where stage = ?", stage)
	if err != nil {
		dao.context.Logger.WithField("getAirdrop", "error").WithField("stage", stage).Errorln("get airdrop by state error")
		return SelectError
	}
	return nil
}
