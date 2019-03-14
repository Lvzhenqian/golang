package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

func addWhiteList(ip string) string {
	args := []string{"/sbin/iptables","-I","INPUT","-s",ip,"-p","tcp","--dport","3879","-j","ACCEPT"}
	cmd := exec.Command("sudo",args...)
	err := cmd.Run()
	if err != nil {
		panic(err.Error())
	}
	return strings.Join(args," ")
}

func getDeleteList() []int {
	var commands []*exec.Cmd
	var ret []int
	args := []string{
		"/sbin/iptables","-nL","--line-num",
	}
	commands = append(commands,exec.Command("sudo",args...))
	commands = append(commands,exec.Command("sudo","grep","3879"))

	commands[1].Stdin,_ = commands[0].StdoutPipe()
	bytes, err := cmd.Output()
	if err != nil {
		//fmt.Println("err!!")
		panic(err)
	}
	fmt.Printf("%s",bytes)
	s := strings.Fields(string(bytes))
	for i,v := range s{
		fmt.Println(i,v)
		ret = append(ret,i)
	}
	return ret
}

func main() {
	//args := WhiteList("192.168.3.22")
	//cmd := exec.Command("sudo" ,args...)
	//b, _ := cmd.Output()
	//fmt.Println(string(b))
	getDeleteList()
}
