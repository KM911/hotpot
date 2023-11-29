package commands

import (
	"fmt"
	"hotpot/config"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func Start() error {
	// sanitize command
	cmd = exec.Command("cmd", "/c", config.UserToml.Command) //nolint:gosec
	if cmd == nil {
		log.Fatal("cmd is nil")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func RunCommand(command string) {
	cmd = exec.Command("cmd", "/c", command)
	cmd.Run()
}

func Stop() error {
	RunCommand("taskkill /F /T /PID " + strconv.Itoa(cmd.Process.Pid))
	fmt.Print("\033[H\033[2J")
	return cmd.Process.Signal(os.Interrupt)

}
