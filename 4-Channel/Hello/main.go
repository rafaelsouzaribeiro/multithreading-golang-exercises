package main

import "fmt"

// thread 1
func main() {
	canal := make(chan string) // canal vazio

	// Thread 2
	go func() {
		canal <- "Hello World" // canal cheio
	}()

	// Thread 1
	msg := <-canal // canal esvazia

	fmt.Println(msg)
}
