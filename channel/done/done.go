package main

import (
	"fmt"
	"sync"
)

func CreateWorker(id int ,wg *sync.WaitGroup) Worker {

	w := Worker{
		make(chan int),
		wg,
	}
	go doWorker(id,w.in,wg)
	return w
}

type Worker struct {
	in 	chan int
	done *sync.WaitGroup
}

func chanDemo() {
	var Workers [10]Worker
	//c := make(chan int)
	var wg sync.WaitGroup
	for i := 0 ;i <10 ; i++ {
		Workers[i] = CreateWorker(i,&wg)
		//go Worker(i,channels[i])
	}
	//wg.Add(20)
	for i,worker := range Workers{
		wg.Add(1)
		worker.in <- 'a' + i
	}

	for i ,worker := range Workers{
		wg.Add(1)
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func doWorker(id int, c chan int, done *sync.WaitGroup) {
	for n := range c{
		fmt.Printf("Worker %d received %c\n", id,n)
		done.Done()
	}
}

func main() {
	//BufferedChannel()
	chanDemo()
	//channelClose()
}
