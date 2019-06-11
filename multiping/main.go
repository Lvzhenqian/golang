package main

import "golang/multiping/ping"

func main() {
	//pinger, err := ping.NewPinger("3.112.70.224",0)
	pinger, err := ping.NewPinger("192.168.3.1",5)
	if err != nil {
		panic(err)
	}
	defer pinger.Close()
	pinger.Run()
}
