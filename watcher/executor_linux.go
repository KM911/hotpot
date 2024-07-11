package watcher

import (
	"os/exec"
)

func CreateCommand(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}

// func Stop() error {
// 	RunCommand("kill " + strconv.Itoa(cmd.Process.Pid))
// 	return cmd.Process.Kill()
// 	// return cmd.Process.Signal(os.Interrupt)
// }
