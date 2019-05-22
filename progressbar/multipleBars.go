package main

import (
	"gopkg.in/cheggaaa/pb.v1"
	"math/rand"
	"sync"
	"time"
)

func main() {
	first := pb.New(200).Prefix("first")
	second := pb.New(200).Prefix("second")
	third := pb.New(200).Prefix("third")

	pool,err := pb.StartPool(first,second,third)
	if err != nil {
		panic(err)
	}
	wg := new(sync.WaitGroup)
	for _,bar := range []*pb.ProgressBar{first,second,third} {
		wg.Add(1)
		go func(cb *pb.ProgressBar) {
			for n:=0;n<200;n++ {
				cb.Increment()
				time.Sleep(time.Millisecond*time.Duration(rand.Intn(100)))
			}
			cb.Finish()
			wg.Done()
		}(bar)
	}
	wg.Wait()
	pool.Stop()
}
