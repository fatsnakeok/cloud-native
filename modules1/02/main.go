package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("producing data: %d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for k := range ch {
		fmt.Printf("Get data: %d\n", k)
	}
}

func main() {
	ch := make(chan int, 6)
	go produce(ch)
	consumer(ch)
}
