package dao

//
//import (
//	"github.com/dometa/bootstrap"
//	"github.com/dometa/model/po"
//)
//
//type userListDao interface {
//	Insert(userList po.UserListPO) (bool, error)
//
//	IsUserJoinAirdrop(address string, airdropTokenId string) (bool, error)
//}
//
//type userListDaoImpl struct {
//	context *bootstrap.Context
//}
//
//func NewUserListDao(context *bootstrap.Context) userListDao {
//	return &userListDaoImpl{
//		context: context,
//	}
//}
//
//// insert user, was make an appointment.
//// make appointment with airdropToken
//func (dao userListDaoImpl) Insert(user po.UserListPO) (bool, error) {
//	insertUserListSql := "insert into userList(walletAddress,airdropTokenId,joinedWhiteList,createTime,updateTime) values(?,?,?,?,?)"
//	res, err := dao.context.MysqlClient.Exec(insertUserListSql, user.WalletAddress, user.AirdropTokenId, user.JoinedWhiteList, user.CreateTime, user.UpdateTime)
//	if err != nil {
//		dao.context.Logger.
//			WithField("insertUserList", "error").
//			WithField("userList", user).Errorln("insert userList error", err)
//		return false, InsertError
//	}
//	affectRows, err := res.RowsAffected()
//	if err != nil {
//		dao.context.Logger.WithField("insertUserList", "noAffectedError").WithField("userList", user).Errorln("no affect error")
//		return false, InsertError
//	}
//	if affectRows > 0 {
//		return true, nil
//	} else {
//		return false, NothingInsertError
//	}
//}
//
//// judgment
//func (dao userListDaoImpl) IsUserJoinAirdrop(walletAddress string, airdropTokenId string) (bool, error) {
//	var userList po.UserListPO
//	sqlStr := "select * from userList where walletAddress=? and airdropTokenId=? and joinedWhiteList=1"
//	err := dao.context.MysqlClient.Get(&userList, sqlStr, walletAddress, airdropTokenId)
//	if err != nil {
//		dao.context.Logger.
//			WithField("selectJoinUser", "error").
//			WithField("walletAddress", walletAddress).
//			WithField("airdropTokenId", airdropTokenId).Errorln("get user has made an appointment", err)
//		return false, err
//	}
//	if userList.Id > 0 {
//		return true, nil
//	} else {
//		return false, NothingSelectError
//	}
//}
