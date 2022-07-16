package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	time "time"
)

var (
	LoggerBuildFailure = errors.New("logger was built in failure.")
	Logger             *log.Logger
	Status             = "Status"
)

var (
	loggerPath                = "logger.path"
	loggerFileSplitInHour     = "logger.fileSplitInHour"
	loggerfileRemoveAfterHour = "logger.fileRemoveAfterHour"
)

func BuildLogger() (*log.Logger, error) {
	fmt.Println("build logger with rolling file")
	logger := log.New()

	logFilePath := viper.GetString(loggerPath)
	splitInHour := viper.GetInt(loggerFileSplitInHour)
	removeAfterHour := viper.GetInt(loggerfileRemoveAfterHour)
	logFileWriter, err := rotatelogs.New(
		//split file format
		logFilePath+"-%Y%m%d%H",
		rotatelogs.WithLinkName(logFilePath),
		rotatelogs.WithMaxAge(time.Duration(splitInHour)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(removeAfterHour)*time.Hour),
	)
	if err != nil {
		fmt.Println("log build throw err => ", err)
	}
	logger.SetOutput(logFileWriter)
	Logger = logger
	return logger, err
}
