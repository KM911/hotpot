package config

const (
	TomlFile = "hotpot.toml"
)

var (
	DefaultToml = Toml{
		Delay:         2000,
		Command:       "go build",
		WatchFiles:    []string{"go"},
		IgnoreFolders: []string{"node_modules", "vendor", ".git", ".idea", ".vscode", "log", "build", "dist", "bin", "public", "target", "output"},
		ShowEvent:     false,
		Github:        "https://github.com/KM911/hotpot",
	}
	UserToml = Toml{}
)

func init() {
	LoadToml()
}
