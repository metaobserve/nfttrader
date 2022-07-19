package controller

import (
	"github.com/devfeel/dotweb"
)

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

// Get banner for homepage.
func GetBanner(ctx dotweb.Context) error {
	return nil
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
