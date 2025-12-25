package udp

import (
	"net"
)

type Notifier struct {
	conn *net.UDPConn
}

func NewNotifier(addr string) *Notifier {
	udpAddr, _ := net.ResolveUDPAddr("udp", addr)
	conn, _ := net.DialUDP("udp", nil, udpAddr)
	return &Notifier{conn: conn}
}

func (n *Notifier) Notify(msg string) {
	_, _ = n.conn.Write([]byte(msg))
}
