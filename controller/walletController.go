package controller

import (
	"github.com/devfeel/dotweb"
	"github.com/dometa/global"
	"github.com/dometa/model"
	"github.com/dometa/model/req"
	"github.com/dometa/model/res"
	"github.com/dometa/service"
	"github.com/dometa/utility"
	"github.com/pkg/errors"
)

var (
	WalletLoginResponseFailure = errors.New("wallet login response failure")
	WalletTokenResponseFailure = errors.New("wallet token response failure")

	walletLogin_Success      = "login success"
	walletLogin_InvalidData  = "login invalid data"
	walletLogin_InvalidToken = "login invalid token"
	walletLogin_Failure      = "login failure"

	WalletToken_Description = "get login token"
)

//WalletLogin
//Add walletAddress
func WalletLogin(ctx dotweb.Context) error {

	//1. invalid data
	result := new(res.WalletLoginResponse)

	walleltLogin := new(req.WalletLoginRequest)

	err := ctx.BindJsonBody(walleltLogin)
	if err != nil {
		global.Logger.WithField("controller", "walletLogin").
			WithField("loginData", "invalid").
			Errorln("invalid walletLogin =>", err)
		result.Description = walletLogin_InvalidData
		result.Status = model.StatusType_CUSTOMERROR
	}

	currentToken, err := global.Cache.Get(walleltLogin.Token)
	if err != nil {
		global.Logger.WithField("controller", "walletLogin").
			WithField("loginToken", "invalid").
			Errorln("invalid login token =>", err)
		result.Description = walletLogin_InvalidToken
		result.Status = model.StatusType_CUSTOMERROR
		return walletLoginResponse(result, ctx)
	}

	if currentToken == nil {
		global.Logger.WithField("controller", "walletLogin").
			WithField("loginToken", "empty").
			Errorln("invalid login token empty =>", err)
		result.Description = walletLogin_InvalidToken
		result.Status = model.StatusType_CUSTOMERROR
		return walletLoginResponse(result, ctx)
	}

	//2. update address, store cookieToken for mind
	cookieToken := walleltLogin.Token
	err = service.WalletLogin(walleltLogin.Address, walleltLogin.Token)
	if err != nil {
		global.Logger.WithField("controller", "walletLogin").
			WithField("loginSave", "failure").
			Errorln("login address save failure =>", err)
		result.Description = walletLogin_Failure
		result.Status = model.StatusType_FAILURE
		return walletLoginResponse(result, ctx)
	}

	//3. response success

	ctx.SetCookieValue(model.Dometa_Cookie, cookieToken, 1*365*24*60*60)
	result.Description = walletLogin_Success
	result.Status = model.StatusType_SUCCESS
	return walletLoginResponse(result, ctx)
}

func walletLoginResponse(result *res.WalletLoginResponse, ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	err := ctx.WriteJson(result)
	if err != nil {
		global.Logger.WithField("controller", "walletLogin").
			WithField("response", "failure").
			Errorln("loginData response Failure =>", err)
		return WalletLoginResponseFailure
	}
	return err
}

func WalletToken(ctx dotweb.Context) error {
	uuid := utility.GetUUID()
	global.Cache.Set(uuid, []byte(uuid))
	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	token := new(res.TokenResponse)
	token.Token = uuid
	token.Description = WalletToken_Description
	token.Status = model.StatusType_SUCCESS
	err := ctx.WriteJson(token)
	if err != nil {
		global.Logger.WithField("controller", "walletToken").
			WithField("response", "failure").
			Errorln("loginToken response Failure =>", err)
		return WalletTokenResponseFailure
	}
	return err
}
