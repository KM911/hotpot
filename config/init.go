package config

const (
	TomlFile = "hotpot.toml"
)

var (
	DefaultToml = Toml{
		Delay:          500,
		PrepareCommand: []string{"go mod tidy", "go build -o bin"}, // block and execute one by one
		HookCommand:    []string{},                                 // not block and execute
		// only one
		Intervals:      1500,
		ExecuteCommand: "./bin",
		WatchFiles:     []string{"*"},
		IgnoreFolders:  []string{"node_modules", "vendor", ".git", ".idea", ".vscode", "log", "build", "dist", "bin", "public", "target", "output"},
		ShowEvent:      true,
		Github:         "https://github.com/KM911/hotpot",
	}
	UserToml = Toml{}
)

func init() {
	UserToml = DefaultToml
}
