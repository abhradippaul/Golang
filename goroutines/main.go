package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println(phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	doneChan <- true
}

func main() {
	go greet("Nice to meed you!")
	go greet("First Hello!")
	done := make(chan bool)
	go slowGreet("How ... are ... you?", done)
	go greet("Second Hello!")
	<-done
}
