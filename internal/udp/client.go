package udp

import (
	"encoding/json"
	"log"
	"net"
)

func Register() {
	{
		conn, err := net.Dial("udp", "localhost:9100")
		if err != nil {
			log.Println("Unable to connect to UDP ")

		}
		msg := Message{Type: "register"}
		data, _ := json.Marshal(msg)

		conn.Write(data)
	}
}
