package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		//m.Unlock()
		atomic.AddUint64(&number, 1)
		time.Sleep(time.Millisecond * 300)
		w.Write([]byte(fmt.Sprintf("você é o visitante %d", number)))
	})
	http.ListenAndServe(":8080", nil)
}
