package dao

import (
	"database/sql"
	"github.com/dometa/global"
	"github.com/dometa/model/po"
)

type WalletDao interface {
	Insert(wallet po.WalletPO) (bool, error)
	Get(wallet *po.WalletPO, address string) error
}

type walletDaoImpl struct {
}

func NewWalletDao() WalletDao {
	return &walletDaoImpl{}
}

func (dao walletDaoImpl) Insert(wallet po.WalletPO) (bool, error) {
	//sqlStr := "insert into wallet(address,token,loginTime,logoutTime,createTime,updateTime)values(?,?,?,?,?,?)"
	sqlStr := "insert into wallet(address,loginTime,logoutTime,createTime)values(?,?,?,?)"
	ret, err := global.MysqlClient.Exec(sqlStr,
		wallet.Address,
		//wallet.Token,
		wallet.LoginTime,
		wallet.LogoutTime,
		wallet.CreateTime,
		//wallet.UpdateTime,
	)
	if err != nil {
		global.Logger.
			WithField("insertWallet", "error").
			WithField("wallet", wallet).Errorln("insert wallet error", err)
		return false, InsertError
	}
	affectRows, err := ret.RowsAffected()
	if err != nil {
		global.Logger.WithField("insertWallet", "noAffectedError").WithField("wallet", wallet).Errorln("no affect error")
		return false, InsertError
	}
	if affectRows > 0 {
		return true, nil
	} else {
		return false, NothingInsertError
	}
}

func (dao walletDaoImpl) Get(wallet *po.WalletPO, address string) error {
	err := global.MysqlClient.Get(wallet, "select id,address from wallet where address = ?", address)
	if err != nil {
		global.Logger.
			WithField("getWallet", "error").
			WithField("address", address).Errorln("get wallet by address error", err)
		if err == sql.ErrNoRows {
			return nil
		}
		return SelectError
	} else {
		return nil
	}
}
