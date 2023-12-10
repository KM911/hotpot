package util

import "runtime"

var (
	KillByname func(name string)
	KillByPid  func(pid string)
)

func killByname(name string) {
	ExecuteCommand("killall " + name)
}

func taskKillByName(name string) {
	ExecuteCommand("taskkill /F /IM " + name)
}

func killByPid(pid string) {
	ExecuteCommand("kill -9 " + pid)
}

func taskKillByPid(pid string) {
	ExecuteCommand("taskkill /F /PID " + pid)
}

func init() {
	if runtime.GOOS == "windows" {
		KillByname = taskKillByName
		KillByPid = taskKillByPid
	} else {
		KillByname = killByname
		KillByPid = killByPid
	}
}
