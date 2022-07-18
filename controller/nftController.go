package controller

import (
	"github.com/devfeel/dotweb"
	"github.com/dometa/dao"
	"github.com/dometa/global"
	"github.com/dometa/model/po"
)

func WalletLogin(ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	//fmt.Println(time.Now(), "Index Handler")
	err := ctx.WriteString("index  => ", ctx.Request().Url())
	return err
}

func Mint(ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	//fmt.Println(time.Now(), "Index Handler")
	err := ctx.WriteString("index  => ", ctx.Request().Url())
	return err
}

func MintAppointment(ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	//fmt.Println(time.Now(), "Index Handler")
	err := ctx.WriteString("index  => ", ctx.Request().Url())
	return err
}

func GetAirdrop(ctx dotweb.Context) error {

	global.Logger.Errorln("testError")
	var airdrop po.AirdropPO
	airportDao := dao.NewAirportDao(global.MysqlClient)
	airportDao.GetAirdropByStage(&airdrop, 1)

	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	//fmt.Println(time.Now(), "Index Handler")
	err := ctx.WriteString("index  => ", ctx.Request().Url())
	return err
}

func GetNfts(ctx dotweb.Context) error {
	return nil
	//values := ctx.Request().URL.Query()
	//var airdropTokenId string
	//airdropTokenId = values.Get("airdropTokenId")
	//pageIndexStr := values.Get("pageIndex")
	//var pageIndex int
	////pageIndex = strconv.ParseInt(pageIndexStr, 8, 16)
	//nftDao := dao.NewNftDao(workContext)
	//var nfts []po.NftPO
	//nftDao.SelectNft(&nfts, airdropTokenId, pageIndex)
	//ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	//if airdropTokenId == nil || pageIndex == nil {
	//	err := ctx.WriteString("[]")
	//	return err
	//}
}
