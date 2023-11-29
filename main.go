package main

import (
	"fmt"
	"github.com/KM911/hotpot/util"
	"github.com/KM911/hotpot/watcher"
	"os"
	"path/filepath"
)

func init() {
	util.FileLogger(filepath.Join(util.ProcessPath, "log.log"))
}

func HelpMessage() {}

func FileModified(filename string) {

}
func main() {
	defer util.Recover(util.ErrorHandler)
	lens := len(os.Args)
	switch lens {
	case 1:
		watcher.StartWatch()
	default:
		fmt.Println("args invalid")
	}
}
