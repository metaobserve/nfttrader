package dao

import (
	"github.com/dometa/bootstrap"
	"github.com/dometa/model/po"
)

type userListDao interface {
	Insert(userList po.UserListPO) (bool, error)

	IsUserJoinAirdrop(address string, airdropTokenId string) (bool, error)
}

type userListDaoImpl struct {
	context *bootstrap.Context
}

func NewUserListDao(context *bootstrap.Context) userListDao {
	return &userListDaoImpl{
		context: context,
	}
}

func (dao userListDaoImpl) Insert(userList po.UserListPO) (bool, error) {
	return false, nil
}

func (dao userListDaoImpl) IsUserJoinAirdrop(address string, airdropTokenId string) (bool, error) {
	var userList po.UserListPO
	sqlStr := "select * from airport where address=? and airdropTokenId=?"
	err := dao.context.MysqlClient.Get(userList, sqlStr, address, airdropTokenId)
	if err != nil {
		return false, err
	}
	if userList.Id > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
