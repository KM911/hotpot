package config

import (
	"fmt"
	"hotpot/util"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Toml struct {
	Delay         int
	Command       string
	WatchFiles    []string
	IgnoreFolders []string
	ShowEvent     bool
	Github        string
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
		//UserToml = DefaultToml
		fmt.Println("Configurations file created, please edit it and restart.")
		os.Exit(0)
	}

	file, err := os.ReadFile(filepath.Join(util.CmdPath(), TomlFile))
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(file, &UserToml)
	if err != nil {
		panic(err)
	}
}
