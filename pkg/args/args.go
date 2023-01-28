package args

import (
	"fmt"
	"os"
)

func CheckCliArgs(configFile *Config) {
	myArgs := map[string]interface{}{"-c": SetConfigFile}
	_, ok := myArgs[os.Args[1]]
	if ok {
		fmt.Println(myArgs[os.Args[1]].(func() string)())
	}
}
