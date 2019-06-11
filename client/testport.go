package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, connerr := net.Dial("tcp","193.112.47.118:22")
	if connerr != nil {
		panic(connerr)
	}

	defer conn.Close()
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		scan.Text()
	}

	fmt.Println(string(buf))
}
