package ping

import (
	"errors"
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"time"
)

var Data = []byte("hello")

type ping struct {
	Addr string
	Conn net.Conn
	Data []byte
	Timeout int
	Count  int
	PacketsSent  int
	PacketsRecv	 int

}

type Reply struct {
	Time  int64
	TTL   uint8
	Error error
}

func NewPinger(addr string,c int) (*ping,error) {
	var count int
	ipaddr, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		return nil, err
	}
	wb, err := MarshalMsg(8, Data)
	if err != nil {
		return nil, err
	}
	if c > 0 {
		count = c
	} else {
		count = 0
	}
	return &ping{Data: wb, Addr: ipaddr.String(),Timeout: 5,Count: count}, nil
}

func MarshalMsg(req int, data []byte) ([]byte, error) {
	xid, xseq := os.Getpid()&0xffff, req
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: xid, Seq: xseq,
			Data: data,
		},
	}
	return wm.Marshal(nil)
}

func (self *ping) Close() error {
	return self.Conn.Close()
}

func (self *ping) Run() {
	var ConnErr error
	self.Conn, ConnErr = net.Dial("ip4:icmp", self.Addr)
	if ConnErr != nil {
		panic(ConnErr)
	}
	self.Conn.SetDeadline(time.Now().Add(time.Duration(self.Timeout) * time.Second))
	c := 0
	for {
		if self.Count > 0 && c >= self.Count {
			return
		}
		r := sendPingMsg(self.Conn, self.Data)
		if r.Error != nil {
			if opt, ok := r.Error.(*net.OpError); ok && opt.Timeout() {
				fmt.Printf("From %s reply: TimeOut\n", self.Addr)
				self.Conn, ConnErr = net.Dial("ip4:icmp", self.Addr)
				if ConnErr != nil {
					return
				}
			} else {
				fmt.Printf("From %s reply: %s\n", self.Addr, r.Error)
			}
		} else {
			fmt.Printf("From %s reply: time=%d ttl=%d\n", self.Addr, r.Time, r.TTL)
		}
		time.Sleep(1e9)
		c++
	}
}

func sendPingMsg(c net.Conn, wb []byte) (reply Reply) {
	start := time.Now()

	if _, reply.Error = c.Write(wb); reply.Error != nil {
		return
	}

	rb := make([]byte, 1500)
	var n int
	n, reply.Error = c.Read(rb)
	if reply.Error != nil {
		return
	}

	duration := time.Since(start)
	ttl := uint8(rb[8])
	rb = func(b []byte) []byte {
		if len(b) < 20 {
			return b
		}
		hdrlen := int(b[0]&0x0f) << 2
		return b[hdrlen:]
	}(rb)
	var rm *icmp.Message
	rm, reply.Error = icmp.ParseMessage(1, rb[:n])
	if reply.Error != nil {
		return
	}

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		t := int64(duration / time.Millisecond)
		reply = Reply{t, ttl, nil}
	case ipv4.ICMPTypeDestinationUnreachable:
		reply.Error = errors.New("Destination Unreachable")
	default:
		reply.Error = fmt.Errorf("Not ICMPTypeEchoReply %v", rm)
	}
	return
}


