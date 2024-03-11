

// 这个文件可以自动生成比较好
package commands
import "github.com/urfave/cli/v2"
var (
	Subcommands []*cli.Command = []*cli.Command{
		&Init,
		&Exec,
		&Watch,
	}
)