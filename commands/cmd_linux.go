package commands

import (
	"fmt"
	"github.com/KM911/hotpot/config"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func Start() error {
	cmd = exec.Command("bash", "-c", config.UserToml.Command) //nolint:gosec
		log.Fatal("cmd is nil")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func RunCommand(command string) {
	cmd = exec.Command("bash", "-c", command)
	cmd.Run()
}

func Stop() error {
	RunCommand("kill -9 " + strconv.Itoa(cmd.Process.Pid))
	//  linux 清空屏幕
	fmt.Println("\033[H\033[2J")
	return cmd.Process.Signal(os.Interrupt)
}
