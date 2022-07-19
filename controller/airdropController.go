package controller

import (
	"github.com/devfeel/dotweb"
	"github.com/dometa/dao"
	"github.com/dometa/global"
	"github.com/dometa/model/po"
)

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
