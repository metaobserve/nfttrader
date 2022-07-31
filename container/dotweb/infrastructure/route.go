package infrastructure

import (
	"github.com/devfeel/dotweb"
	"github.com/dometa/controller"
)

func InitRoute(container *dotweb.DotWeb) {
	//container.HttpServer.Router().GET("/banner", controller.GetBanner)

	container.HttpServer.Router().GET("/wallet/token", controller.WalletToken)
	container.HttpServer.Router().POST("/wallet/login", controller.WalletLogin)

	//container.HttpServer.Router().GET("/nft/airdrop", controller.GetAirdrop)

	container.HttpServer.Router().POST("/nft/page", controller.Nfts)
	container.HttpServer.Router().POST("/nft/mint", controller.Mint)
	//container.HttpServer.Router().POST("/nft/mintAppointment", controller.MintAppointment)
}
