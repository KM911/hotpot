package config

const (
	TomlFile = "hotpot.toml"
)

var (
	DefaultToml = Toml{
		Delay:         500,
		Command:       "go run main.go",
		WatchFiles:    []string{"*"},
		IgnoreFolders: []string{"node_modules", "vendor", ".git", ".idea", ".vscode", "log", "build", "dist", "bin", "public", "target", "output"},
		ShowEvent:     true,
		Github:        "https://github.com/KM911/hotpot",
	}
	UserToml = Toml{}
)

func init() {
	UserToml = DefaultToml
}
