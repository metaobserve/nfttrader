package dotweb

import (
	"fmt"
	web "github.com/devfeel/dotweb"
	"github.com/dometa/container/dotweb/infrastructure"
	"github.com/dometa/container/dotweb/internal"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	containerLoggerDirectory = "application.container.loggerDirectory"
	containerPort            = "application.container.port"
)
var (
	WebContainerBuildError = errors.New("dotweb build error")
)

func StartServer() error {

	loggerDirectory := viper.GetString(containerLoggerDirectory)
	port := viper.GetInt(containerPort)

	container := web.New()
	container.SetLogPath(loggerDirectory)
	container.HttpServer.SetEnabledBindUseJsonTag(true)
	container.HttpServer.SetEnabledSession(true)
	container.HttpServer.SetEnabledIgnoreFavicon(true)

	container.HttpServer.SetBinder(internal.NewTypeBinder())
	infrastructure.InitRoute(container)
	fmt.Println("dometa server start")
	err := container.StartServer(port)
	if err != nil {
		fmt.Println("dometa server start failure =>", err)
		return WebContainerBuildError
	}

	fmt.Println("dometa server start success => :", port)

	return nil
}
