package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/format"
	"github.com/KM911/hotpot/watcher"

	"github.com/KM911/hotpot/util"
)

var (
	FunctionMatch = map[string]func(){}
)

func init() {
	format.FileLogger(filepath.Join(util.ExecuteDirectory, "log.log"))
}

func HelpMessage() {}

// flag parse and start watch

// hotpot init
// hotpot watch
// hotpot help

// 我希望hotpot可以支持在命令行启动时 直接传入命令
// 如果没有configuration 就用默认的配置
// 反之就用配置文件的配置 只不过command会被覆盖
func main() {
	lens := len(os.Args)
	switch lens {
	case 2:
		if os.Args[1] == "watch" || os.Args[1] == "w" {
			watcher.StartWatch()
		} else {
			config.UserToml.Command = strings.Join(os.Args[1:], "")
			watcher.StartWatch()
			fmt.Println("args invalid")
		}
	default:
		watcher.StartWatch()
	}

}
