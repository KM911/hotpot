package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/KM911/fish/format"
	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/watcher"
	"github.com/urfave/cli/v2"
)

var (
	Exec = cli.Command{
		Name:    "exec",
		Usage:   "Use argv as command and watch",
		Aliases: []string{"e"},
		Action:  ExecAction,
	}
)

func ExecAction(c *cli.Context) error {
	config.UserToml.ExecuteCommand = strings.Join(c.Args().Slice(), " ")
	starship, err := exec.LookPath("starship")
	format.Must(err)

	if starship != "" {
		for l, ext := range Languages {
			r, w, _ := os.Pipe()
			defer w.Close()
			guess := exec.Cmd{
				Path:   starship,
				Args:   []string{"", "module", l},
				Stdout: w,
				Stderr: w,
			}
			guess.Run()
			if guess.ProcessState.ExitCode() == 0 {
				r.Close()
				config.UserToml.WatchFiles = append(config.UserToml.WatchFiles, ext)
				break
			}

		}
	}
	watcher.ProcessWatchEnvironment()
	watcher.StartWatch()
	return nil
}

var (
	Languages = map[string]string{
		"golang":     "go",
		"java":       "java",
		"javascript": "js",
		"ruby":       "rb",
		"php":        "php",
		"csharp":     "cs",
		"cpp":        "cpp",
		"python":     "py",
		"rust":       "rs",
	}
)
