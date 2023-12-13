package main

import "fmt"

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("hello", hello)
	ler(hello)
}

// Apenas vai receber nosso canal (chan<- string)
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// Somente para entregar resultado esvaziar(<-chan string)
func ler(data <-chan string) {
	fmt.Println(<-data)
}
