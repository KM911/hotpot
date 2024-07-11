package watcher

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/KM911/fish/format"

	"github.com/KM911/hotpot/config"
)

var (
	Start func() error
	Count = 0
)

func Stop() error {
	return cmd.Process.Kill()
}

func StartBindHook() error {
	Count++
	format.Info("Execute times" + strconv.Itoa(Count))
	for _, command := range config.UserToml.PrepareCommand {
		RunCommand(command)
	}
	time.Sleep(time.Duration(config.UserToml.Delay) * time.Millisecond)
	cmd = CreateCommand(config.UserToml.ExecuteCommand)
	if cmd == nil {
		log.Fatal("cmd is nil")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	time.Sleep(time.Duration(config.UserToml.Intervals) * time.Millisecond)
	SocketPing()
	return nil
}
func StartWithHook() error {
	Count++
	format.Info("Execute times" + strconv.Itoa(Count))
	for _, command := range config.UserToml.PrepareCommand {
		RunCommand(command)
	}
	StopHookCommand()
	go RunHookCommand(config.UserToml.HookCommand)
	time.Sleep(time.Duration(config.UserToml.Intervals) * time.Millisecond)
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

func init() {
	if config.HookEnable {
		Start = StartBindHook

	} else {
		Start = StartWithHook
	}
}
