package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	inPutReader := bufio.NewReader(os.Stdin)
	fmt.Println("first, what is your name? ")
	clientName, _ := inPutReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	for {
		fmt.Println("what to send to the server? Type Q to quit.")
		input, _ := inPutReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if strings.ToLower(trimmedInput) == "q" {
			return
		}
		_, err := conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		if err != nil {
			fmt.Println("Send Error", err.Error())
			continue
		}
	}
}
