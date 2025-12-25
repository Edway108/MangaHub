package main

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"

	"MangaHub/internal/udp"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":9091")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("UDP server on :9091")

	clients := make(map[string]*net.UDPAddr)
	var mu sync.Mutex

	buf := make([]byte, 2048)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		var msg udp.Notification
		_ = json.Unmarshal(buf[:n], &msg)

		switch msg.Type {

		case "register":
			mu.Lock()
			clients[clientAddr.String()] = clientAddr
			mu.Unlock()
			log.Println("Client registered:", clientAddr.String())

		case "broadcast":
			log.Println("Broadcast:", msg.Message)

			mu.Lock()
			for _, c := range clients {
				out, _ := json.Marshal(udp.Notification{
					Type:      "notify",
					Message:   msg.Message,
					Timestamp: time.Now().Unix(),
				})
				conn.WriteToUDP(out, c)
			}
			mu.Unlock()
		}
	}
}
