package commands

import (
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/KM911/hotpot/util"

	"github.com/KM911/hotpot/config"
)

var (
	Stop func() error
)

func Start() error {
	cmd = util.CreateCommand(config.UserToml.Command)
	if cmd == nil {
		log.Fatal("cmd is nil")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func RunCommand(command string) error {
	cmd = util.CreateCommand(command)
	return cmd.Run()
}

func taskkill() error {
	println("\033[H\033[2J")
	RunCommand("taskkill /F /T /PID " + strconv.Itoa(cmd.Process.Pid))
	return cmd.Process.Signal(os.Interrupt)

}

func kill() error {
	println("\033[H\033[2J")
	RunCommand("kill -9 " + strconv.Itoa(cmd.Process.Pid))
	return cmd.Process.Signal(os.Interrupt)
}

func init() {
	if runtime.GOOS == "windows" {
		Stop = taskkill
	} else {
		Stop = kill
	}
}
