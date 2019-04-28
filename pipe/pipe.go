package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

func UnixPipe() {
	cmd1 := exec.Command("ls","-l")
	cmd2 := exec.Command("wc","-l")
	stdout1,_ := cmd1.StdoutPipe()
	if err:=cmd1.Start();err != nil {
		fmt.Printf("Error: can't runing cmd1 %s\n",err)
		return
	}
	outPutBuf1 := bufio.NewReader(stdout1)
	stdin2,_ := cmd2.StdinPipe()
	outPutBuf1.WriteTo(stdin2)
	var outPutBuf2 bytes.Buffer
	cmd2.Stdout = &outPutBuf2
	if err := cmd2.Start();err != nil{
		fmt.Printf("Error: can't running cmd2 %s\n",err)
		return
	}
	if err := stdin2.Close(); err != nil {
		fmt.Printf("Error: can't close cmd2 stdinIO %s\n",err)
		return
	}

	if err:= cmd2.Wait();err != nil{
		fmt.Printf("Error: Can not wait for the command: %s\n", err)
		return
	}
	fmt.Printf("%s\n",outPutBuf2.Bytes())
}


func main() {
	UnixPipe()
}
