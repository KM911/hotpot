package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// 这部分应该剥离开来的 不然太二了 不是吗
const (
	// 应该是一个
	DefaultTomlFilename = "config_km911.toml"
)

// func CreateDefaultTomlConfig() {
// 	CreateTomlConfig(filepath.Join("~", ".config", DefaultTomlFilename), DefaultToml)
// }

func CreateTomlConfig(_file string, _toml interface{}) {
	f, err := os.Create(_file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = toml.NewEncoder(f).Encode(_toml)
	if err != nil {
		panic(err)
	}
}

func SaveTomlConfig(_file string, _toml interface{}) {
	// os.WriteFile(_file, []byte(toml.Marshal(_toml)), 0777 )
	f, err := os.OpenFile(_file, os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	toml.NewEncoder(f).Encode(_toml)

}

func LoadTomlConfig(_file string, _toml interface{}) {
	if _, err := os.Stat(_file); os.IsNotExist(err) {
		CreateTomlConfig(_file, _toml)
		fmt.Println("Configs file created, please edit it and restart.")
		os.Exit(0)
	}
	// 是不是应该解耦呢? 就是如果修改为 ~/.config/config 就可以溢出这里的依赖来不是吗
	// cmd path or execute path
	file, err := os.ReadFile(_file)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(file, _toml)
	if err != nil {
		panic(err)
	}

}

func LoadDefaultTomlConfig() {
	// LoadTomlConfig(filepath.Join("~", ".config", DefaultTomlFilename))

	// if _, err := os.Stat(DefaultTomlFilename); os.IsNotExist(err) {
	// 	CreateDefaultTomlConfig()
	// 	fmt.Println("Configs file created, please edit it and restart.")
	// 	os.Exit(0)
	// }
	// // 是不是应该解耦呢? 就是如果修改为 ~/.config/config 就可以溢出这里的依赖来不是吗
	// // cmd path or execute path
	// file, err := os.ReadFile(filepath.Join("~", ".config", DefaultTomlFilename))
	// if err != nil {
	// 	panic(err)
	// }
	// err = toml.Unmarshal(file, &UserToml)
	// if err != nil {
	// 	panic(err)
	// }
}
