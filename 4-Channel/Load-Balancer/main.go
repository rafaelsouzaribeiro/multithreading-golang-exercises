package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qntWorker := 100000
	// Multiplos workers para ajudar na executa(Ganho de performace)
	for i := 0; i < qntWorker; i++ {
		// Está lendo data toda vez que é modficada
		go worker(i, data)
	}

	// Ele tá enviando dados para data
	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
