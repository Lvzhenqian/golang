package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	go SourceData("t.txt")
	go Tailf("t.txt")
	for i:=0;i<60;i++{
		//fmt.Printf("main %d\n",i)
		time.Sleep(1*1e9)
	}
}

func SourceData(FileName string)  {
	buf, err := os.OpenFile(FileName, os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer buf.Close()
	OutWrite := bufio.NewWriter(buf)
	n := 0
	for {
		//if n == 10 {
		//	break
		//}
		var str string
		str = "now is "+ strconv.Itoa(n) + "\r\n"
		//fmt.Print(str)
		_,err := OutWrite.WriteString(str)
		if err != nil {
			return
		}
		OutWrite.Flush()
		n++
		time.Sleep(1*1e9)
	}
}

func Tailf(Filename string)  {
	buf, err := os.OpenFile(Filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("ERR: %s",err.Error())
		return
	}
	defer buf.Close()
	OutRead := bufio.NewReader(buf)
	for {
		line,err := OutRead.ReadString('\n')
		if err != nil{
			continue
		}
		fmt.Printf("%s",line)
	}

}