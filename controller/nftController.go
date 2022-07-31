package controller

import (
	"github.com/devfeel/dotweb"
	"github.com/dometa/global"
	"github.com/dometa/model"
	"github.com/dometa/model/req"
	"github.com/dometa/model/res"
	"github.com/dometa/service"
	"github.com/pkg/errors"
)

var (
	NftsResponseFailure = errors.New("nfts response failure")

	nfts_Success      = "nfts reqeust success"
	nfts_InvalidData  = "nfts reqeust invalid data"
	nfts_InvalidToken = "nfts reqeust invalid token"
	nfts_Failure      = "nfts reqeust failure"
)

//func Mint(ctx dotweb.Context) error {
//	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
//	//fmt.Println(time.Now(), "Index Handler")
//	err := ctx.WriteString("index  => ", ctx.Request().Url())
//	return err
//}
//
//func MintAppointment(ctx dotweb.Context) error {
//	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
//	//fmt.Println(time.Now(), "Index Handler")
//	err := ctx.WriteString("index  => ", ctx.Request().Url())
//	return err
//}

//// Get banner for homepage.
//func GetBanner(ctx dotweb.Context) error {
//	return nil
//}

func Nfts(ctx dotweb.Context) error {
	//1. invalid data
	response := new(res.NftResponse)
	request := new(req.NftRequest)
	ctx.BindJsonBody(request)
	//2. mint
	nftService := service.NewNftService()

	var nfts []res.NftItem
	hasMoreNfts, err := nftService.GetNfts(&nfts, request.Category, request.PageIndex)

	if err != nil {
		global.Logger.WithField("controller", "mint").
			WithField("nftsRequest", "failure").
			Errorln("nftsRequest  failure =>", err)
		response.Description = nfts_Failure
		response.Status = model.StatusType_FAILURE
		return nftsResponse(response, ctx)
	}

	response.Nfts = nfts
	response.HasMoreNFT = hasMoreNfts

	//3. response success
	response.Description = nfts_Success
	response.Status = model.StatusType_SUCCESS
	return nftsResponse(response, ctx)
}

func nftsResponse(result *res.NftResponse, ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	err := ctx.WriteJson(result)
	if err != nil {
		global.Logger.WithField("controller", "nfts").
			WithField("response", "failure").
			Errorln("nfts response Failure =>", err)
		return NftsResponseFailure
	}
	return err
}
