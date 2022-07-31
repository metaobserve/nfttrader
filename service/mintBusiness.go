package service

import (
	"github.com/dometa/dao"
	"github.com/dometa/global"
	"github.com/dometa/model/po"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

var (
	Mint_InvalidAddress      = errors.New("mint invalid address")
	Mint_SelectError         = errors.New("select address error before mint")
	Mint_SaveFailure         = errors.New("mint save address failure")
	Mint_CheckFailure        = errors.New("mint check failure")
	Mint_AirdropCheckFailure = errors.New("mint airdrop check failure")
	Mint_NoAirdropAction     = errors.New("no mint airdrop action")
	Mint_InsertFailure       = errors.New("mint add failure")

	// airdrop is playing.
	AirDrop_ACTIONStage = 1
	nftRandomRange      = "nft.randomRange"
	nftRandomStartIndex = "nft.randomStartIndex"
)

func Mint(address string) (string, error) {

	if address == "" && len(address) < 15 {
		return "invalid wallet address", Mint_InvalidAddress
	}

	// 1. whether login before
	walletDao := dao.NewWalletDao()
	walletPO := new(po.WalletPO)
	err := walletDao.Get(walletPO, address)
	if err != nil {
		global.Logger.WithField("business", "mint").
			WithField("selectAddress", "error").Errorln("select wallet address error => ", err)
		return "wallet address fail found", WalletLogin_SelectError
	}

	if walletPO.Id < 1 {
		global.Logger.WithField("business", "mint").Infoln("wallet address not exist")
		return "wallet address not found", WalletLogin_SelectError
	}

	// 2.0 check mint bfore ?
	// whether mint before
	mintPO := new(po.MintPO)
	mintDao := dao.NewMintDao()
	err = mintDao.Get(mintPO, address)
	if err != nil {
		global.Logger.WithField("business", "mint").Infoln("mint check failure ==>", err)
		return "mint check filuare", Mint_CheckFailure
	}
	if mintPO.Id > 1 {
		return "had mint before", nil
	}

	// 3.0 mint
	// random airdrop nft
	nftStartIndex := viper.GetInt(nftRandomStartIndex)
	nftRange := viper.GetInt(nftRandomRange)
	nftIndex := nftStartIndex + rand.Intn(nftRange)

	mintPO.Address = address
	// find nftId
	mintPO.NftId = nftIndex
	// find airDropId
	airportDao := dao.NewAirportDao()
	airDropPO := new(po.AirdropPO)
	err = airportDao.GetAirdropByStage(airDropPO, AirDrop_ACTIONStage)

	if err != nil {
		global.Logger.WithField("business", "mint").Infoln("airdrop check failure ==>", err)
		return "airdrop check failure", Mint_AirdropCheckFailure
	}
	if airDropPO.Id < 1 {
		global.Logger.WithField("business", "mint").Infoln("no airdrop action")
		return "no airdrop action", Mint_NoAirdropAction
	}

	mintPO.AirdropId = airDropPO.Id
	mintPO.CreateTime.Time = time.Now()
	mintPO.CreateTime.Valid = true

	mintPO.UpdateTime.Time = time.Now()
	mintPO.UpdateTime.Valid = true

	// manual manager before
	// mint add new nft
	mintPO.NftId = nftIndex

	ret, err := mintDao.Insert(*mintPO)
	if err != nil {
		global.Logger.WithField("business", "mint").Infoln("mint add error")
		return "mint add filuare", Mint_InsertFailure
	}
	if ret == false {
		global.Logger.WithField("business", "mint").Infoln("mint add failure")
		return "mint add filuare", Mint_InsertFailure
	}

	global.Logger.WithField("business", "mint").Infoln("mint success")
	return "mint success", nil
}
