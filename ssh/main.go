package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

func main() {
	session,err := connetc("root","","",22)
	if err != nil{
		log.Fatal(err)
	}
	defer session.Close()
	cmd := "/bin/bash /root/a.sh"


	reader,err := session.StdoutPipe()
	if err != nil { log.Fatal(err)}
	scanner :=bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	session.Run(cmd)

}


func connetc(user string, password string, host string, port int) (*ssh.Session, error) {
	var (
		auth			[]ssh.AuthMethod
		addr			string
		clientConfig	*ssh.ClientConfig
		client			*ssh.Client
		session			*ssh.Session
		err				error
	)
	auth = make([]ssh.AuthMethod,0)
	auth = append(auth,ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User: 		user,
		Auth:		auth,
		Timeout:	30*time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	addr = fmt.Sprintf("%s:%d",host,port)
	if client,err = ssh.Dial("tcp",addr,clientConfig); err != nil{
		return nil,err
	}
	if session, err = client.NewSession();err!=nil{
		return nil,err
	}
	return session,nil
}
