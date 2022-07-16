package dotweb

import (
	"fmt"
	web "github.com/devfeel/dotweb"
	"github.com/dometa/bootstrap"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	containerLoggerDirectory = "application.container.loggerDirectory"
	containerPort            = "application.container.port"
)
var (
	WebContainerBuildFailure = errors.New("dotweb build failure")
)

func BuildWebContainer(ctx *bootstrap.Context) (*bootstrap.Context, error) {

	loggerDirectory := viper.GetString(containerLoggerDirectory)
	port := viper.GetInt(containerPort)

	container := web.New()
	container.SetLogPath(loggerDirectory)
	err := container.StartServer(port)
	if err != nil {
		fmt.Println("dotweb build failure =>", err)
		return ctx, WebContainerBuildFailure
	}

	ctx.WebContainer = container
	return ctx, nil
}
