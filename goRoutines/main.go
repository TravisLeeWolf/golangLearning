package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet("friend, nice to meet you", done)
	go greet(" other friend, how are you?", done)
	go slowGreet("I...see..you...finally...made...your...appearance.", done)
	go greet("my good sir, can I offer you a cup of coffee?", done)

	for range done {
	}
}

func greet(phrase string, doneChannel chan bool) {
	fmt.Println("Hello", phrase)
	doneChannel <- true
}

func slowGreet(phrase string, doneChannel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello there", phrase)
	doneChannel <- true
	close(doneChannel)
}
