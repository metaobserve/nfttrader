package dao

import (
	"github.com/dometa/bootstrap"
	"github.com/dometa/model/po"
)

type WalletDao interface {
	Insert(wallet po.WalletPO) (bool, error)
	Get(wallet *po.WalletPO, address string) error
}

type walletDaoImpl struct {
	context *bootstrap.Context
}

func NewWalletDao(context *bootstrap.Context) WalletDao {
	return &walletDaoImpl{
		context: context,
	}
}

func (dao walletDaoImpl) Insert(wallet po.WalletPO) (bool, error) {
	sqlStr := "insert into wallet(address,token,loginTime,createTime,updateTime)values(?,?,?,?,?)"
	ret, err := dao.context.MysqlClient.Exec(sqlStr,
		wallet.Address,
		wallet.Token,
		wallet.LoginTime,
		wallet.CreateTime,
		wallet.UpdateTime,
	)
	if err != nil {
		dao.context.Logger.
			WithField("insertWallet", "error").
			WithField("wallet", wallet).Errorln("insert wallet error", err)
		return false, FailInsertError
	}
	affectRows, err := ret.RowsAffected()
	if err != nil {
		dao.context.Logger.WithField("insertWallet", "noAffectedError").WithField("wallet", wallet).Errorln("no affect error")
		return false, FailInsertError
	}
	if affectRows > 0 {
		return true, nil
	} else {
		return false, NoneInsertError
	}
}

func (dao walletDaoImpl) Get(wallet *po.WalletPO, address string) error {
	err := dao.context.MysqlClient.Get(wallet, "select * from wallet where address = ?", address)
	if err != nil {
		dao.context.Logger.
			WithField("getWallet", "error").
			WithField("address", address).Errorln("get wallet by address error", err)
		return SelectError
	} else {
		return nil
	}
}
