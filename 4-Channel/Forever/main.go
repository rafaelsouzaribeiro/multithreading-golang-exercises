package main

// thread 1
func main() {
	forever := make(chan bool) // canal vazio

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true // cheio
	}()

	//<- forever = true (Não vai funcionar porque ele precisa rodar em uma go func como encima)

	<-forever // esvaziar
}
