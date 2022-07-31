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
	MintResponseFailure = errors.New("mint response failure")

	mint_Success      = "mint success"
	mint_InvalidData  = "mint invalid data"
	mint_InvalidToken = "mint invalid token"
	mint_Failure      = "mint failure"
)

func DoMetaMint(ctx dotweb.Context) error {

	//1. invalid data
	response := new(res.MintResponse)
	mint := new(req.MintRequest)

	address, err := ctx.Request().Cookie(model.Dometa_Cookie)
	if err != nil {
		global.Logger.WithField("controller", "mint").
			WithField("address", "read from cookie").
			Errorln("read address failure =>", err)
		response.Description = mint_InvalidData
		response.Status = model.StatusType_CUSTOMERROR
		return mintResponse(response, ctx)
	}
	if address == nil || len(address.Value) < 15 {
		global.Logger.WithField("controller", "mint").
			Errorln("invalid wallet address")
		response.Description = mint_InvalidData
		response.Status = model.StatusType_CUSTOMERROR
		return mintResponse(response, ctx)
	}
	mint.Address = address.Value

	//2. mint
	ret, err := service.Mint(mint.Address)

	if err != nil {
		global.Logger.WithField("controller", "mint").
			WithField("mintSave", "failure").
			Errorln("mint address save failure =>", err)
		response.Description = mint_Failure
		response.Status = model.StatusType_FAILURE
		return mintResponse(response, ctx)
	}

	//3. response success
	response.Description = ret
	response.Status = model.StatusType_SUCCESS
	return mintResponse(response, ctx)
}

func mintResponse(result *res.MintResponse, ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	err := ctx.WriteJson(result)
	if err != nil {
		global.Logger.WithField("controller", "mint").
			WithField("response", "failure").
			Errorln("mint response Failure =>", err)
		return MintResponseFailure
	}
	return err
}

//TODO: had mint
