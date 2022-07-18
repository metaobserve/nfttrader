package infrastructure

import (
	"github.com/devfeel/dotweb"
	"github.com/dometa/controller"
)

func InitRoute(container *dotweb.DotWeb) {
	container.HttpServer.Router().GET("/nft/show", controller.GetNfts)
	container.HttpServer.Router().GET("/nft/airdrop", controller.GetAirdrop)

	container.HttpServer.Router().POST("/nft/walletLogin", controller.WalletLogin)
	container.HttpServer.Router().POST("/nft/mint", controller.Mint)
	container.HttpServer.Router().POST("/nft/mintAppointment", controller.MintAppointment)
}
