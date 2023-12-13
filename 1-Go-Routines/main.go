package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
	}
}

// Thread 1 (Main)
func main() {
	//Thread 2
	go task("A")
	//Thread 3
	go task("B")

	// func anonima thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonimos")
			time.Sleep(time.Second * 1)
		}
	}()
	// nada sai por isso o sleep

	time.Sleep(time.Second * 15)
}
