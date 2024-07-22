package main

import (
	_ "net/http/pprof"

	"github.com/KM911/hotpot/app"
)

func main() {

	// go func() {
	// 	http.ListenAndServe("0.0.0.0:8080", nil)
	// }()
	app.NewApp("hotpot", "watch file change and execute command")
}
