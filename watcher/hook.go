package watcher

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/KM911/fish/format"
	"github.com/KM911/hotpot/config"
)

var (
	conn       net.Conn
	SocketPing func()
)

func ConnectSocket() {
	conn, err = net.Dial("unix", config.HotpotSocketAddress)
	if err != nil {
		fmt.Println("connect error:", err)
	}
	fmt.Println("connect success")
}

func SocketPingAction() {
	conn.Write([]byte("ping"))
}

func TempHookNotify() {
	file, err := os.OpenFile(filepath.Join(os.TempDir(), "hotpot.sock"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	format.Must(err)
	_, err = file.WriteString("1")
	defer file.Close()
	format.Must(err)
}
func SocketPingEmpty() {}
