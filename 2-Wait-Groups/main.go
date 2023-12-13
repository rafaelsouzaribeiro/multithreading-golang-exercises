package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, w *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
		// cada vez que passar pelo DONE() tira uma operação
		w.Done()
	}
}

// Thread 1 (Main)
func main() {
	waitgroup := sync.WaitGroup{}
	// 25 operações
	waitgroup.Add(25)
	//Thread 2
	go task("A", &waitgroup)
	//Thread 3
	go task("B", &waitgroup)

	// func anonima thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonimos")
			time.Sleep(time.Second * 1)
			waitgroup.Done()
		}
	}()
	waitgroup.Wait()
}
