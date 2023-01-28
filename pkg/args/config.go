package args

import (
	"errors"
	"fmt"
	"os"
)

func SetConfigFile() string {
	filepath := os.Args[2]

	if len(filepath) == 0 {
		panic("No file specified")
	}
	if _, err := os.OpenFile(filepath, os.O_RDONLY, 0755); errors.Is(err, os.ErrNotExist) || errors.Is(err, os.ErrPermission) {
		panic(fmt.Sprintf("%s or %s", os.ErrNotExist, os.ErrPermission))
	}
	return filepath
}
