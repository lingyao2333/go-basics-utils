package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/lingyao2333/go-basics-utils/watching"
)

func init() {
	http.HandleFunc("/lockorder1", lockorder1)
	http.HandleFunc("/lockorder2", lockorder2)
	http.HandleFunc("/req", req)
	go http.ListenAndServe(":10003", nil)
}

func main() {
	w := watching.NewWatching(
		watching.WithCollectInterval("5s"),
		watching.WithCoolDown("1m"),
		watching.WithDumpPath("/tmp"),
		watching.WithTextDump(),
		watching.WithGoroutineDump(10, 25, 2000, 10000),
	)
	w.EnableGoroutineDump().Start()
	time.Sleep(time.Hour)
}

var (
	l1 sync.Mutex
	l2 sync.Mutex
)

func req(wr http.ResponseWriter, req *http.Request) {
	l1.Lock()
	defer l1.Unlock()
}

func lockorder1(wr http.ResponseWriter, req *http.Request) {
	l1.Lock()
	defer l1.Unlock()

	time.Sleep(time.Minute)

	l2.Lock()
	defer l2.Unlock()
}

func lockorder2(wr http.ResponseWriter, req *http.Request) {
	l2.Lock()
	defer l2.Unlock()

	time.Sleep(time.Minute)

	l1.Lock()
	defer l1.Unlock()
}
