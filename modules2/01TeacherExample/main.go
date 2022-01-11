package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var fmeng = false

func producer(threadID int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for !fmeng {
		time.Sleep(1 * time.Second)
		count++
		//strconv.Itoa函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字。
		data := strconv.Itoa(threadID) + "-----" + strconv.Itoa(count)
		fmt.Printf("producer, %s\n", data)

		ch <- data
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("consumer, %s\n", data)
	}
	wg.Done()
}

func main() {

	chanStream := make(chan string, 10)

	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanStream)
	}

	for i := 0; i < 2; i++ {
		wgCs.Add(1)
		go consumer(wgCs, chanStream)
	}

	go func() {
		time.Sleep(time.Second * 3)
		fmeng = true
	}()

	wgPd.Wait()

	close(chanStream)
	wgCs.Wait()

}
