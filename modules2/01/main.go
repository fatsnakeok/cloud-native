package main

/**
只写通道：chan<- T
只读通道：<-chan T
*/
import (
	"fmt"
	"time"
)

func consumer(cname string, ch <-chan int) {

	for i := range ch {
		fmt.Println("consumer-----", cname, ":", i)
	}

	fmt.Println("ch closed")
}

func producer(pname string, ch chan<- int) {

	for i := 0; i < 4; i++ {
		fmt.Println("prooducer----", pname, ":", i)
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	go producer("生产者1", ch)
	go producer("生产者2", ch)
	go consumer("消费者1", ch)
	go consumer("消费者2", ch)

	time.Sleep(10 * time.Second)
	close(ch)
	time.Sleep(10 * time.Second)
}
