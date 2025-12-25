package udp

import (
	"encoding/json"
	"net"
	"time"
)

func Broadcast(message string) error {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:9091")
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	msg := Notification{
		Type:      "broadcast",
		Message:   message,
		Timestamp: time.Now().Unix(),
	}

	data, _ := json.Marshal(msg)
	_, err = conn.Write(data)
	return err
}
