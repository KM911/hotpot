package watcher

import (
	"fmt"
	"github.com/KM911/hotpot/config"
	"net"
)

var (
	conn net.Conn
)

func ConnectSocket() {
	conn, err = net.Dial("unix", config.HotpotSocketAddress)
	if err != nil {
		fmt.Println("connect error:", err)
	}
	fmt.Println("connect success")
}

func SocketPing() {
	conn.Write([]byte("ping"))
}
