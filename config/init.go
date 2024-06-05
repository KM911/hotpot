package config

import (
	"os"
	"path/filepath"

	"github.com/KM911/hotpot/lib/util"
)

var (
	TomlFile string
)

var (
	HookEnable  = false
	DefaultToml = ConfigToml{
		Delay:          500,
		PrepareCommand: []string{"go mod tidy", "go build -o bin"}, // block and execute one by one
		HookCommand:    "",                                         // not block and execute
		Intervals:      0,
		ExecuteCommand: "./bin",
		WatchFiles:     []string{"go"},
		IgnoreFolders:  []string{"node_modules", "vendor", ".git", ".idea", ".vscode", "log", "build", "dist", "bin", "public", "target", "output"},
		ShowEvent:      true,
		Github:         "https://github.com/KM911/hotpot",
	}
	UserToml = ConfigToml{}
)

func init() {
	UserToml = DefaultToml
	// check /tmp/hotpot.sock
	if _, err := os.Stat(HotpotSocketAddress); err == nil {
		HookEnable = true
	}
	TomlFile = filepath.Join(util.WorkingDirectory, "hotpot.toml")
}
