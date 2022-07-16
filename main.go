package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/dometa/bootstrap"
	"github.com/dometa/model/po"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func main() {

	context, err := bootstrap.Start()
	context.Logger.Infoln("hello world")

	var accountPOs []po.AccountPO
	err = context.MysqlClient.Select(&accountPOs, "select * from account")
	fmt.Println("dotweb error =>", err)
	//fmt.Printf(accountPOs[0])
	for i := 0; i < len(accountPOs); i++ {
		fmt.Println(accountPOs[i].Id)
	}
	fmt.Println("hello world")
	app := dotweb.New()

	app.SetLogPath("/home/logs/wwwroot/")
	app.HttpServer.GET("/coco.png", func(ctx dotweb.Context) error {
		return ctx.File("coco.png")
		//return ctx.Attachment("coco.png", "coco.png")
		//return ctx.WriteString("welcome to my first web")
	})

	viper.SetDefault("test", "test")

	fmt.Println(viper.GetString("test"))

	fmt.Println("web begin")
	//err := app.StartServer(8033)
	//fmt.Println("dotweb error =>", err)

	var files []string

	root := "./utility"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println("dotweb error =>", err)
}
