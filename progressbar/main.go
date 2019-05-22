package main

import (
	"gopkg.in/cheggaaa/pb.v1"
	"time"
)

func main() {
	count := 10000
	bar := pb.StartNew(count)
	for i:=0;i<count;i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
}
