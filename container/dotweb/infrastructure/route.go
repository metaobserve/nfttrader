package infrastructure

import (
	"github.com/dometa/bootstrap"
	"github.com/dometa/controller"
)

func InitRoute(context *bootstrap.Context) {

	controller.InitController(context)

	context.WebContainer.HttpServer.Router().GET("/nft/show", controller.GetNfts)
	context.WebContainer.HttpServer.Router().GET("/nft/airdrop", controller.GetAirdrop)

	context.WebContainer.HttpServer.Router().POST("/nft/walletLogin", controller.WalletLogin)
	context.WebContainer.HttpServer.Router().POST("/nft/mint", controller.Mint)
	context.WebContainer.HttpServer.Router().POST("/nft/mintAppointment", controller.MintAppointment)
}
