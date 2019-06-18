package main

import (
	"fmt"
	"golang/multiping/ping"
	"sync"
	"time"
)



func worker(ip string, c int, wg *sync.WaitGroup) {
	wk, err :=ping.NewPinger(ip,c)
	if err != nil {
		panic(err)
	}
	wk.Interval = time.Second * 1
	wk.Timeout = time.Second *	1
	wk.OnRecv = func(r ping.Reply) {
		fmt.Printf("from %s icmp_seq=%d ttl=%d time=%v\n",r.Addr,r.Seq,r.TTL,r.Time)
	}

	wk.OnTimeOut = func(r ping.Reply) {
		fmt.Printf("from %s icmp_seq=%d timeout !!\n",r.Addr,r.Seq)
	}
	defer wk.Close()
	wk.Run()
	st := wk.Getstatistics()
	fmt.Printf("%s addr, %d send, %d recv, %v%% loss\n",st.Addr,st.SendPackets,st.RecvPackets,st.LostPercent)
	wg.Done()
}


func main() {
	ips := []string{"192.168.8.1","115.231.9.79"}
	var wg sync.WaitGroup
	for _,v := range ips{
		wg.Add(1)
		go worker(v,5,&wg)
	}
	wg.Wait()
}