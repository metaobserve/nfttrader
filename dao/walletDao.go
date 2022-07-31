package dao

import (
	"database/sql"
	"github.com/dometa/global"
	"github.com/dometa/model/po"
	"github.com/pkg/errors"
)

var (
	Wallet_UpdateError = errors.New("update wallet error")
)

type WalletDao interface {
	Insert(wallet po.WalletPO) (bool, error)
	GetByToken(wallet *po.WalletPO, token string) error
	GetByAddress(wallet *po.WalletPO, address string) error
	UpdateWalltToken(address string, token string) (bool, error)
}

type walletDaoImpl struct {
}

func NewWalletDao() WalletDao {
	return &walletDaoImpl{}
}

func (dao walletDaoImpl) Insert(wallet po.WalletPO) (bool, error) {
	//sqlStr := "insert into wallet(address,token,loginTime,logoutTime,createTime,updateTime)values(?,?,?,?,?,?)"
	sqlStr := "insert into wallet(address,token,loginTime,logoutTime,createTime)values(?,?,?,?)"
	ret, err := global.MysqlClient.Exec(sqlStr,
		wallet.Address,
		wallet.Token,
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

func (dao walletDaoImpl) GetByToken(wallet *po.WalletPO, token string) error {
	err := global.MysqlClient.Get(wallet, "select id,address,token from wallet where token = ?", token)
	if err != nil {
		global.Logger.
			WithField("getWallet", "error").
			WithField("token", token).Errorln("get wallet by token error", err)
		if err == sql.ErrNoRows {
			return nil
		}
		return SelectError
	} else {
		return nil
	}
}

func (dao walletDaoImpl) GetByAddress(wallet *po.WalletPO, address string) error {
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

func (dao walletDaoImpl) UpdateWalltToken(address string, token string) (bool, error) {
	ret, err := global.MysqlClient.Exec("update wallet set token=? where address = ?", token, address)
	if err != nil {
		global.Logger.
			WithField("updateWalletToken", "error").
			WithField("address", address).Errorln("get wallet by address error ==>", err)
		return false, Wallet_UpdateError
	}

	affectRows, err := ret.RowsAffected()
	if err != nil {
		global.Logger.WithField("updateWalletToken", "error").Errorln("update token no affetced ==>", err)
		return false, Wallet_UpdateError
	}
	if affectRows > 0 {
		return true, nil
	} else {
		return false, NothingInsertError
	}
}
