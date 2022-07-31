package bootstrap

import (
	"fmt"
	container "github.com/dometa/container/dotweb"
)

func Start() error {

	err := startWebServer()
	if err != nil {
		return err
	}
	return nil
}

func startWebServer() error {

	err := container.StartServer()
	if err != nil {
		fmt.Println("build webContainer failure =>", err)
		return nil
	}
	fmt.Println("build web container success")
	return nil
}
