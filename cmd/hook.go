package cmd

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/KM911/hotpot/config"
	"github.com/KM911/hotpot/watcher"
	"github.com/fsnotify/fsnotify"

	"github.com/KM911/fish/format"
	"github.com/KM911/fish/system"
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
func UnixWatcheServer() {
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

}
func HookAction(c *cli.Context) error {
	config.LoadToml()
	watcher.ProcessWatchEnvironment()
	//
	TempHookAction()
	return nil
}

func TempHookNotify() {
	// append data into file
	file, err := os.OpenFile(filepath.Join(os.TempDir(), "hotpot.sock"), os.O_CREATE|os.O_WRONLY, 0666)
	format.Must(err)
	defer file.Close()
	file.WriteString("1")
}
func TempHookAction() {
	file, err := os.Create(filepath.Join(os.TempDir(), "hotpot.sock"))
	format.Must(err)
	defer file.Close()

	watcher, err := fsnotify.NewWatcher()
	format.Must(err)
	defer watcher.Close()
	err = watcher.Add(filepath.Join(os.TempDir(), "hotpot.sock"))
	format.Must(err)
	pc := 0
	for _ = range watcher.Events {
		println("\033[H\033[2J")
		pc++
		format.InfoMessage("Times :", strconv.Itoa(pc))
		system.ExecuteCommand(config.UserToml.HookCommand)
	}
	fmt.Println("exit temphook action")
}
