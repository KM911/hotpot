package watcher

import (
	"fmt"
	"net"

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

func SocketPingEmpty() {}
