package dao

import (
	"github.com/dometa/bootstrap"
	"github.com/dometa/model/po"
)

type whiteListDao interface {
	Insert(whiteList po.WhiteListPO) (bool, error)

	IsUserInWhiteList(airdropTokenId string, walletAddress string) (bool, error)
}

type whiteListDaoImpl struct {
	context *bootstrap.Context
}

func NewWhiteListDao(context *bootstrap.Context) whiteListDao {
	return &whiteListDaoImpl{
		context: context,
	}
}

// insert user, was make an appointment.
// make appointment with airdropToken
func (dao whiteListDaoImpl) Insert(whiteList po.WhiteListPO) (bool, error) {

	insertUserListSql := "insert into whiteList(name,description,walletAddress,airdropTokenId,startTime,endTime,status,createTime,updateTime) values(?,?,?,?,?,?,?,?,?)"
	ret, err := dao.context.MysqlClient.Exec(insertUserListSql,
		whiteList.Name, whiteList.Description, whiteList.WalletAddress,
		whiteList.AirdropTokenId, whiteList.StartTime, whiteList.EndTime, whiteList.Status, whiteList.CreateTime, whiteList.UpdateTime)
	if err != nil {
		dao.context.Logger.
			WithField("insertWhiteList", "error").
			WithField("whiteList", whiteList).Errorln("insert whiteList error", err)
		return false, InsertError
	}
	affectRows, err := ret.RowsAffected()
	if err != nil {
		dao.context.Logger.WithField("insertWhiteList", "noAffectedError").WithField("whiteList", whiteList).Errorln("no affect error")
		return false, InsertError
	}
	if affectRows > 0 {
		return true, nil
	} else {
		return false, NothingInsertError
	}
}

// judgment
func (dao whiteListDaoImpl) IsUserInWhiteList(airdropTokenId string, walletAddress string) (bool, error) {
	var whiteList po.WhiteListPO
	sqlStr := "select * from whiteList where airdropTokenId=? and walletAddress=? and status=1"
	err := dao.context.MysqlClient.Get(&whiteList, sqlStr, airdropTokenId, walletAddress)
	if err != nil {
		dao.context.Logger.
			WithField("findUserInWhiteList", "error").
			WithField("airdropTokenId", airdropTokenId).
			WithField("walletAddress", walletAddress).Errorln("find user in whiteList", err)
		return false, err
	}
	if whiteList.Id > 0 {
		return true, nil
	} else {
		return false, NothingSelectError
	}
}
