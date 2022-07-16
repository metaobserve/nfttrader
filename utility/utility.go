package utility

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

var (
	GetExecutePathFailure = errors.New("get execute path failure")
)

func GetCurrentExecutePath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("get execute path failure =>", err)
		return EmptyString, GetExecutePathFailure
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}
