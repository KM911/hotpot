package commands

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/KM911/hotpot/lib/format"

	"github.com/KM911/hotpot/config"
)

var (
	Stop  func() error
	Count = 0
)

func CreateCommand(command string) *exec.Cmd {
	return exec.Command("sh", "-c", command)

}

func Start() error {
	Count++
	format.InfoMessage("Execute times", strconv.Itoa(Count))
	cmd = CreateCommand(config.UserToml.Command)
	if cmd == nil {
		log.Fatal("cmd is nil")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func RunCommand(command string) error {
	cmd = CreateCommand(command)
	return cmd.Run()
}

func taskkill() error {

	RunCommand("taskkill /F /T /PID " + strconv.Itoa(cmd.Process.Pid))
	return cmd.Process.Signal(os.Interrupt)

}

func kill() error {
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
