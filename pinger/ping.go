package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"os"
	"os/signal"
)


func main() {
	pinger ,err := ping.NewPinger("www.baidu.com")
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal,1)
	signal.Notify(c,os.Interrupt)
	pinger.SetPrivileged(true)

	go func() {
		for _ = range c{
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",pkt.Nbytes,pkt.IPAddr,pkt.Seq,pkt.Rtt)
	}

	pinger.OnFinish = func(stat *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n",stat.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stat.PacketsSent,stat.PacketsRecv,stat.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stat.MinRtt,stat.AvgRtt,stat.MaxRtt,stat.StdDevRtt)
	}
	fmt.Printf("Ping %s (%s):\n",pinger.Addr(),pinger.IPAddr())
	pinger.Run()
}