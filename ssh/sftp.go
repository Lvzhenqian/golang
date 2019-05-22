package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net"
	"strings"
	"time"
)

func listdir(ip string,port int,username string,password string,Src string) error {
	ClientCf := &ssh.ClientConfig{
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		User: username,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client,clienterr := ssh.Dial("tcp", fmt.Sprintf("%s:%d",ip,port), ClientCf)
	if clienterr != nil {
		return clienterr
	}
	defer client.Close()
	// new SftpClient
	sftpClient, err := sftp.NewClient(client)
	defer sftpClient.Close()
	if err != nil {
		return err
	}
	walk:=sftpClient.Walk(Src)
	for walk.Step() {
		var Prefix string = ""
		if walk.Stat().IsDir(){
			Prefix = "dir"
		} else {
			Prefix = "file"
		}
		s:=strings.Join([]string{Prefix,walk.Path()},":")
		fmt.Println(s)
	}
	return nil
}

func main() {
	listdir("193.112.47.118",22,"root","charles90!@","/data/jupyter")
}