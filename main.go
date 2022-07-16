package main

import (
	"fmt"
	"github.com/dometa/bootstrap"
	"github.com/dometa/dao"
	"github.com/dometa/model/po"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {

	context, err := bootstrap.Start()
	if err != nil {
		fmt.Println("dometa start failure =>", nil)
	} else {
		fmt.Println("dometa started.")
	}

	walletDao := dao.NewWalletDao(context)
	wallet := new(po.WalletPO)
	wallet.Token = "abc"
	wallet.Address = "abc"
	wallet.LoginTime.Time = time.Now()
	wallet.CreateTime.Time = time.Now()
	wallet.UpdateTime.Time = time.Now()

	var wallet2 po.WalletPO
	//walletDao.Insert(*wallet)
	err = walletDao.Get(&wallet2, "abc")
	fmt.Println(wallet2)
	fmt.Println(err)

	//StartTime   string       `db:"startTime"`
	//EndTime     sql.NullTime `db:"endTime"`
	//// airDrop stage: 0 ready; 1 action; 2 cut
	//Stage      int          `db:"stage"`
	//CreateTime sql.NullTime `db:"createTime"`
	//UpdateTime sql.NullTime `db:"updateTime"`
	var airdrop po.AirdropPO
	airdrop.Name = "天下第一空投"
	airdrop.Description = "这是我们的第一次空头"
	airdrop.StartTime.Time = time.Now()
	airdrop.EndTime.Time = time.Now()
	airdrop.CreateTime.Time = time.Now()
	airdrop.Stage = 1

	//airportDao := dao.NewAirportDao(context)
	//airportDao
	//var accountPOs []po.AccountPO
	//err = context.MysqlClient.Select(&accountPOs, "select * from account")
	//fmt.Println("dotweb error =>", err)
	////fmt.Printf(accountPOs[0])
	//for i := 0; i < len(accountPOs); i++ {
	//	fmt.Println(accountPOs[i].Id)
	//}
	//fmt.Println("hello world")
	//app := dotweb.New()
	//
	//app.SetLogPath("/home/logs/wwwroot/")
	//app.HttpServer.GET("/coco.png", func(ctx dotweb.Context) error {
	//	return ctx.File("coco.png")
	//	//return ctx.Attachment("coco.png", "coco.png")
	//	//return ctx.WriteString("welcome to my first web")
	//})
	//
	//viper.SetDefault("test", "test")
	//
	//fmt.Println(viper.GetString("test"))
	//
	//fmt.Println("web begin")
	////err := app.StartServer(8033)
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
