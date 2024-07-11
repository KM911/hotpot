package watcher

import (
	"os/exec"
)

func CreateCommand(command string) *exec.Cmd {

	return exec.Command("cmd", "/c", command)
}

// func Stop() error {

// 	RunCommand("taskkill /F /T /PID " + strconv.Itoa(cmd.Process.Pid))
// 	return cmd.Process.Signal(os.Interrupt)

// }
