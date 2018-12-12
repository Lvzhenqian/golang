package ssh

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
)

func Test() {
	ce := func(err error,msg string) {
		if err != nil {
			log.Fatalf("%s error: %v",msg,err)
		}
	}

	client, err := ssh.Dial("tcp","192.168.8.6:22",&ssh.ClientConfig{
		User: "root",Auth: []ssh.AuthMethod{ssh.Password("zw-9898w")},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	ce(err,"dial")
	session,err := client.NewSession()
	ce(err,"new session")
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:	0,
		ssh.TTY_OP_ISPEED:	14400,
		ssh.TTY_OP_OSPEED:	14400,
	}
	err = session.RequestPty("linux",25,80,modes)
	ce(err,"request pty")
	err = session.Shell()
	ce(err,"start shell")
	err = session.Wait()
	ce(err,"return")
}