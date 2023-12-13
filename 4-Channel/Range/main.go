package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int)
	w := sync.WaitGroup{}
	w.Add(10)

	go publish(ch)
	go reader(ch, &w)

	// wait group segura o processamento
	w.Wait()
}

// lê as mensagens que está publicando no canal
func reader(ch chan int, w *sync.WaitGroup) {
	for v := range ch {
		fmt.Printf("Received %d\n", v)
		w.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// fechar o canal nunca mais vai publicar
	close(ch)
}
