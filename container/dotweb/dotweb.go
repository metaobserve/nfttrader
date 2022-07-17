package dotweb

import (
	"fmt"
	web "github.com/devfeel/dotweb"
	"github.com/dometa/container/dotweb/infrastructure"
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

func BuildWebContainer() (*web.DotWeb, error) {

	loggerDirectory := viper.GetString(containerLoggerDirectory)
	port := viper.GetInt(containerPort)

	container := web.New()
	container.SetLogPath(loggerDirectory)
	container.HttpServer.SetEnabledBindUseJsonTag(true)

	infrastructure.InitRoute(container)

	err := container.StartServer(port)
	if err != nil {
		fmt.Println("dotweb build failure =>", err)
		return nil, WebContainerBuildError
	}

	fmt.Println("dotweb start => :", port)

	return container, nil
}
