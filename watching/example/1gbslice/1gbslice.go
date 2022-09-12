package main

import (
	"net/http"
	"time"

	"github.com/lingyao2333/go-basics-utils/watching"
)

func init() {
	http.HandleFunc("/make1gb", make1gbslice)
	go http.ListenAndServe(":10003", nil)
}

func main() {
	w := watching.NewWatching(
		watching.WithCollectInterval("2s"),
		watching.WithCoolDown("1m"),
		// watching.WithDumpPath("./tmp"),
		watching.WithTextDump(),
		watching.WithMemDump(1, 5, 10),
	)
	w.EnableMemDump().Start()
	time.Sleep(time.Hour)
}

func make1gbslice(wr http.ResponseWriter, req *http.Request) {
	a := make([]byte, 1073741824)
	_ = a
}
