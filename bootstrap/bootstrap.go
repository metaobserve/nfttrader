package bootstrap

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/dometa/component/database/mysql"
	"github.com/dometa/component/logger"
	"github.com/dometa/component/resources"
	"github.com/dometa/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Context struct {
	Logger       *log.Logger
	MysqlClient  *sqlx.DB
	WebContainer *dotweb.DotWeb
}

func Start() (*Context, error) {
	fmt.Println("bootstrap: application initialize")
	context := new(Context)
	fmt.Println("bootstrap: read environment resources")
	err := buildEnvResources()
	if err != nil {
		return context, err
	}

	context, err = buildLogger(context)
	if err != nil {
		return context, err
	}

	context, err = buildMySqlClient(context)
	if err != nil {
		return context, err
	}
	return context, nil
}

// build environment first
func buildEnvResources() error {
	err := resources.BuildEnvResources()
	if err != nil {
		return resources.ReadConfigurationFailure
	}
	return nil
}

func buildLogger(context *Context) (*Context, error) {
	fmt.Println("build a logger and inject into all the layers.")
	l, err := logger.BuildLogger()
	if err != nil {
		fmt.Println("build logger failure")
		return context, logger.LoggerBuildFailure
	}
	context.Logger = l
	context.Logger.
		WithField("build", "logger").
		WithField(logger.Status, model.StatusType_SUCCESS).
		Infoln("build mysql client success")
	return context, nil
}

func buildWebContainer(context *Context) *Context {
	fmt.Println("build the webContainer")
	return context
}

func buildMySqlClient(context *Context) (*Context, error) {
	fmt.Println("build myql client")
	client, err := mysql.BuildMySqlClient()
	if err != nil {
		context.Logger.
			WithField("build", "mysqlClient").
			WithField(logger.Status, model.StatusType_EXCEPTION).Errorln("build mysql client failure", err)
		fmt.Println("build mysql client failure =>", err)
	}
	context.MysqlClient = client
	context.Logger.
		WithField("build", "mysqlClient").
		WithField(logger.Status, model.StatusType_SUCCESS).
		Infoln("build mysql client success")
	return context, nil
}

func Dispose(ctx *Context) {
	fmt.Println("dispose all resources")
	if ctx != nil {
		ctx.Logger.WithField("dispose", "resources").Infoln("dispose all resources.")
		ctx.Logger.WithField("webContainer", "close").Infoln("close dotweb container.")
		ctx.WebContainer.Close()
		ctx.Logger.WithField("mysqlClient", "close").Infoln("close mysql connections.")
		ctx.MysqlClient.Close()
	}

}
