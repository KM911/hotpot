package main

import (
	"fmt"
	"github.com/KM911/hotpot/format"
	"os"
	"path/filepath"

	"github.com/KM911/hotpot/util"
	"github.com/KM911/hotpot/watcher"
)

func init() {

	format.FileLogger(filepath.Join(util.ExecuteDirectory, "log.log"))
}

func HelpMessage() {}

func FileModified(filename string) {

}
func main() {

	lens := len(os.Args)
	switch lens {
	case 1:
		watcher.StartWatch()
	default:
		fmt.Println("args invalid")
	}
}
