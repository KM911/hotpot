package config

import (
	"os"
	"path/filepath"

	"github.com/KM911/fish/fs"
)

var (
	TomlFile string
)

var (
	HookEnable  = false
	DefaultToml = ConfigToml{
		Delay:          500,
		PrepareCommand: []string{"go mod tidy"}, // block and execute one by one
		HookCommand:    "",                      // not block and execute
		Intervals:      0,
		ExecuteCommand: "go run main.go",
		WatchFiles:     []string{"go"},
		IgnoreFolders:  []string{"node_modules", "vendor", ".git", ".idea", ".vscode", "log", "build", "dist", "bin", "public", "target", "output"},
		ShowEvent:      true,
		Github:         "https://github.com/KM911/hotpot",
	}
	UserToml = DefaultToml
)

func init() {
	// check /tmp/hotpot.sock
	if _, err := os.Stat(HotpotSocketAddress); err == nil {
		HookEnable = true
	}
	TomlFile = filepath.Join(fs.WorkingDirectory, "hotpot.toml")
}
