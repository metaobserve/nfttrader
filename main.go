package main

import (
	"fmt"
	"github.com/dometa/bootstrap"
	"github.com/dometa/core"
	"github.com/dometa/global"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	core.BuildEnvResources()
	logger, err := core.BuildLogger()

	global.Logger = logger

	mySqlClient, err := core.BuildMySqlClient()
	global.MysqlClient = mySqlClient

	err = bootstrap.Start()
	if err != nil {
		fmt.Println("dometa start failure =>", err)
	} else {
		fmt.Println("dometa started.")
	}

	//airdrop

	//var airdrop po.AirdropPO
	//airdrop.Name = "天下第一空投"
	//airdrop.Description = "这是我们的第一次空头"
	//airdrop.StartTime.Time = time.Now()
	//airdrop.StartTime.Valid = true
	//airdrop.EndTime.Time = time.Now()
	//airdrop.EndTime.Valid = true
	//airdrop.CreateTime.Time = time.Now()
	//
	//airdrop.UpdateTime.Time = time.Now()
	//airdrop.UpdateTime.Valid = true
	//airdrop.Stage = 1
	//airdrop.TokenId = utility.GetGUID()
	//
	//airportDao := dao.NewAirportDao(context)
	//ret, err := airportDao.Insert(airdrop)
	//var airdropPO po.AirdropPO
	//airportDao.GetAirdropByStage(&airdropPO, 1)
	//fmt.Println(ret)
	//airportDao

	//nft
	//var nft po.NftPO
	//nft.Token = "token"
	//nft.Name = "name"
	//nft.Status = 1
	//nft.Description = "description"
	//nft.CreateTime.Valid = true
	//nft.CreateTime.Time = time.Now()
	//nft.UpdateTime.Valid = true
	//nft.UpdateTime.Time = time.Now()
	//nft.AirdropTokenId = "token"
	//nftDao := dao.NewNftDao(context)
	//ret, err := nftDao.Insert(nft)
	//fmt.Println(ret)
	//fmt.Println(err)
	//
	//var nfts []po.NftPO
	//err = nftDao.SelectNft(&nfts, nft.AirdropTokenId)
	//fmt.Println(err)

	//userList
	//var user po.UserListPO
	//user.WalletAddress = "address"
	//user.AirdropTokenId = "token"
	//user.JoinedWhiteList = true
	//user.CreateTime.Valid = true
	//user.CreateTime.Time = time.Now()
	//user.UpdateTime.Valid = true
	//user.UpdateTime.Time = time.Now()
	//
	//userListDao := dao.NewUserListDao(context)
	//ret, err := userListDao.Insert(user)
	//fmt.Println("ret =>", ret)
	//fmt.Println("err =>", err)
	//
	//ret, err = userListDao.IsUserJoinAirdrop("address", "token")
	//fmt.Println("ret =>", ret)
	//fmt.Println("err =>", err)

	//whiteList
	//var whiteList po.WhiteListPO
	//whiteList.AirdropTokenId = "token"
	//whiteList.WalletAddress = "address"
	//whiteList.Name = "name"
	//whiteList.Description = "description"
	//whiteList.Status = 1
	//whiteList.StartTime.Valid = true
	//whiteList.StartTime.Time = time.Now()
	//whiteList.EndTime.Valid = true
	//whiteList.EndTime.Time = time.Now()
	//
	//whiteListDao := dao.NewWhiteListDao(context)
	//ret, err := whiteListDao.Insert(whiteList)
	//fmt.Println("ret =>", ret)
	//fmt.Println("err =>", err)
	//
	//ret, err = whiteListDao.IsUserInWhiteList("token", "address")
	//fmt.Println("ret =>", ret)
	//fmt.Println("err =>", err)

	// wallet
	//walletDao := dao.NewWalletDao(context)
	//var wallet po.WalletPO
	//wallet.Token = "abc"
	//wallet.Address = "abc"
	//wallet.LoginTime.Time = time.Now()
	//wallet.LoginTime.Valid = true
	//
	//wallet.CreateTime.Time = time.Now()
	//wallet.CreateTime.Valid = true
	//
	//wallet.UpdateTime.Time = time.Now()
	//wallet.UpdateTime.Valid = true
	//
	//wallet.LogoutTime.Time = time.Now()
	//wallet.LogoutTime.Valid = true
	//
	//walletDao.Insert(wallet)
	//var wallet2 po.WalletPO
	//fmt.Println(wallet2.Id)
	//err = walletDao.Get(&wallet2, "abc")
	//fmt.Println(wallet2)
	//fmt.Println(err)

	//context.WebContainer.HttpServer.GET("/coco.png", func(ctx dotweb.Context) error {
	//	return ctx.File("coco.png")
	//	//return ctx.Attachment("coco.png", "coco.png")
	//	//return ctx.WriteString("welcome to my first web")
	//})
	//
	//err = context.WebContainer.StartServer(8033)
	//fmt.Println("err =>", err)

	////fmt.Println("dotweb error =>", err)
	//
	//var files []string
	//
	//root := "./utility"
	//err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	//	files = append(files, path)
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//for _, file := range files {
	//	fmt.Println(file)
	//}
	//fmt.Println("dotweb error =>", err)
}
