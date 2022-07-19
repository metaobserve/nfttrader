package controller

import (
	"github.com/devfeel/dotweb"
)

func WalletLogin(ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	//fmt.Println(time.Now(), "Index Handler")
	err := ctx.WriteString("index  => ", ctx.Request().Url())
	return err
}
