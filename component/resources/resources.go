package resources

import (
	"fmt"
	"github.com/dometa/utility"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	ReadConfigurationFailure = errors.New("read application configuration failure")
)

func BuildEnvResources() error {
	path, err := utility.GetCurrentExecutePath()
	if err != nil {
		fmt.Println("get execute path failure => ", err)
		return err
	}

	viper.AddConfigPath(path)
	viper.AddConfigPath(path + "/resources")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./resources")
	viper.SetConfigName("application")
	//TODO: split configurations into file.

	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("read configuration err =>", err)
		return ReadConfigurationFailure
	}
	return nil
}
