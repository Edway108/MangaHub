package udp

import (
	"encoding/json"
	"log"
	"net"
)

func StartClient() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:9999")
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	
	register := Notification{
		Type: "register",
	}
	data, _ := json.Marshal(register)
	conn.Write(data)

	log.Println("UDP client registered")

	// 🔹 LISTEN
	buf := make([]byte, 2048)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("UDP read error:", err)
			continue
		}

		var msg Notification
		if err := json.Unmarshal(buf[:n], &msg); err != nil {
			continue
		}

		log.Println(
			"[UDP NOTIFY]",
			msg.MangaID,
			msg.Message,
		)
	}
}

func SendNotification(content string) error {
	addr, err := net.ResolveUDPAddr("udp", "localhost:9999")
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	msg := Message{
		Type:    "notify",
		Content: content,
	}

	data, _ := json.Marshal(msg)
	_, err = conn.Write(data)
	return err
}
