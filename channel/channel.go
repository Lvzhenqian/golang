package main

import (
	"fmt"
	"time"
)

func CreateWorker(id int ) chan<- int {
	c := make(chan  int)
	go worker(id,c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	//c := make(chan int)
	for i := 0 ;i <10 ; i++ {
		channels[i] = CreateWorker(i)
		//go worker(i,channels[i])
	}

	for i := 0; i < 10; i++{
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++{
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
	for n := range c{
		fmt.Printf("Worker %d received %c\n", id,n)
	}
}

func BufferedChannel() {
	c := make(chan int,3)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	go worker(0,c)
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int,3)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	go worker(0,c)
	close(c)
	time.Sleep(time.Millisecond)
}


func main() {
	//BufferedChannel()
	//chanDemo()
	channelClose()
}
