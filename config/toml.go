package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KM911/fish/fs"
	"github.com/pelletier/go-toml/v2"
)

var (
	HotpotSocketAddress string
)

func init() {
	HotpotSocketAddress = filepath.Join(os.TempDir(), strings.ReplaceAll(fs.WorkingDirectory, "/", "_")+"hotpot.sock")
}

type ConfigToml struct {
	Delay          int
	PrepareCommand []string
	HookCommand    string
	Intervals      int
	ExecuteCommand string
	WatchFiles     []string
	IgnoreFolders  []string
	ShowEvent      bool
	Github         string
}

func CreateDefaultToml() {
	f, err := os.Create(TomlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = toml.NewEncoder(f).Encode(DefaultToml)
	if err != nil {
		panic(err)
	}
}

func LoadToml() {
	if _, err := os.Stat(TomlFile); os.IsNotExist(err) {
		CreateDefaultToml()
		fmt.Println("Configurations file created, please edit it and restart.")
		os.Exit(0)
	}
	file, err := os.ReadFile(TomlFile)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(file, &UserToml)
	if err != nil {
		panic(err)
	}
}
