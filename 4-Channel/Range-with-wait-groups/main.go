package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

// lê as mensagens que está publicando no canal
func reader(ch chan int) {
	for v := range ch {
		fmt.Printf("Received %d\n", v)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// fechar o canal nunca mais vai publicar
	close(ch)
}
