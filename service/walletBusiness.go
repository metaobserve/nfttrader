package service

import (
	"github.com/dometa/dao"
	"github.com/dometa/global"
	"github.com/dometa/model/po"
	"github.com/pkg/errors"
	"time"
)

var (
	WalletLogin_InvalidAddress = errors.New("select address error before login address")
	WalletLogin_SelectError    = errors.New("select address error before login address")
	WalletLogin_SaveFailure    = errors.New("save address failure")
)

func WalletLogin(address string, token string) error {

	if address == "" && len(address) < 15 {
		return WalletLogin_InvalidAddress
	}

	walletDao := dao.NewWalletDao()
	walletPO := new(po.WalletPO)
	err := walletDao.GetByAddress(walletPO, address)
	if err != nil {
		global.Logger.WithField("business", "walletLogin").
			WithField("selectAddress", "error").Errorln("select Address error => ", err)
		return WalletLogin_SelectError
	}

	if walletPO.Id > 0 {
		walletDao.UpdateWalltToken(address, token)
		global.Logger.WithField("business", "walletLogin").Infoln("wallet address had login")
		return nil
	}

	walletPO.Address = address
	walletPO.Token = token
	walletPO.CreateTime.Valid = true
	walletPO.LoginTime.Time = time.Now()
	walletPO.LoginTime.Valid = true

	walletPO.CreateTime.Time = time.Now()
	walletPO.CreateTime.Valid = true

	walletPO.UpdateTime.Time = time.Now()
	walletPO.UpdateTime.Valid = true
	ret, err := walletDao.Insert(*walletPO)

	if err != nil || ret == false {
		global.Logger.WithField("business", "walletLogin").
			WithField("saveAddress", "failure").Errorln("save Address failure => ", err)
		return WalletLogin_SaveFailure
	}
	global.Logger.WithField("business", "walletLogin").Infoln("login success")
	return nil
}
