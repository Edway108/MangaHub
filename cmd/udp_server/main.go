package main

import (
	"MangaHub/internal/udp"
	"encoding/json"
	"fmt"
	"net"
)

func main() {

	addr, _ := net.ResolveUDPAddr("udp", ":9100")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("UDP Notification Server running on :9100")

	clients := make(map[string]*net.UDPAddr)
	buf := make([]byte, 1024)

	for {
		n, ClientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		var msg udp.Message
		json.Unmarshal(buf[:n], &msg)

		if msg.Type == "register" {
			clients[ClientAddr.String()] = ClientAddr
			fmt.Println("UDP client registered:", ClientAddr)

			if msg.Type == "broadcast" {
				data, _ := json.Marshal(msg)
				for _, addr := range clients {
					conn.WriteToUDP(data, addr)

				}
			}

		}

	}
}
