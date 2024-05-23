package watcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/KM911/hotpot/lib/format"

	"github.com/KM911/hotpot/config"
)

var (
	Stop  func() error
	Count = 0
)

func CreateCommand(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)

}

func Start() error {
	Count++
	format.InfoMessage("Execute times", strconv.Itoa(Count))

	// CreateCommand(config.UserToml.BuildCommand).Run()
	for _, command := range config.UserToml.PrepareCommand {
		RunCommand(command)
	}
	StopHookCommand()
	for _, command := range config.UserToml.HookCommand {
		go RunHookCommand(command)
	}
	time.Sleep(time.Duration(config.UserToml.Delay))
	cmd = CreateCommand(config.UserToml.ExecuteCommand)
	if cmd == nil {
		log.Fatal("cmd is nil")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func RunCommand(command string) error {
	cmd = CreateCommand(command)
	err := cmd.Run()
	fmt.Println("RunCommand", err)
	return err
}
func RunHookCommand(command string) error {
	cmd := CreateCommand(command)
	err := cmd.Run()
	hookCmds <- cmd
	return err
}

func StopHookCommand() {
	for {
		select {
		case cmd := <-hookCmds:
			cmd.Process.Signal(os.Interrupt)
		default:
			return
		}
	}
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
