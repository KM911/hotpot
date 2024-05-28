package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KM911/hotpot/lib/util"

	"github.com/pelletier/go-toml/v2"
)

const (
	HotpotSocketAddress = "/tmp/hotpot.sock"
)

type Toml struct {
	Delay          int
	PrepareCommand []string
	HookCommand    string
	Intervals      int
	ExecuteCommand string
	WatchFiles     []string
	IgnoreFolders  []string
	EnableHook     bool
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
	file, err := os.ReadFile(filepath.Join(util.WorkingDirectory, TomlFile))
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(file, &UserToml)
	if err != nil {
		panic(err)
	}
}
