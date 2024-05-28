package commands

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/lib/format"
	"github.com/KM911/hotpot/watcher"
	"github.com/urfave/cli/v2"
)

var (
	Hook = cli.Command{
		Name:    "hook",
		Usage:   "hook command",
		Aliases: []string{"h"},
		Action:  HookAction,
	}
	pc = 0
)

func echoServer(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal("read error:", err)
			}
			break
		}
		// system.ExecuteCommand(config.UserToml.HookCommand)
		// fmt.Println("Received:", string(buf[:n]))
		// watcher.RunCommand(config.UserToml.HookCommand)
		println("\033[H\033[2J")
		format.InfoMessage("Execute times", strconv.Itoa(pc))
		pc++
		cmd := watcher.CreateCommand(config.UserToml.HookCommand)
		if cmd == nil {
			log.Fatal("cmd is nil")
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()

	}
}

func HookAction(c *cli.Context) error {
	config.LoadToml()

	if !config.UserToml.EnableHook {
		fmt.Println("Please Enable Hook in the configuration file")
		return nil
	}
	watcher.ProcessWatchEnvironment()
	os.RemoveAll(config.HotpotSocketAddress)
	l, err := net.Listen("unix", config.HotpotSocketAddress)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(conn)
	}
	return nil
}
