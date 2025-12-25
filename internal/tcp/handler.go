package tcp

import (
	"MangaHub/pkg/utils"
	"encoding/json"
	"log"
	"net"

	"github.com/google/uuid"
)
func HandleConnection(conn net.Conn, hub *Hub) {
	defer conn.Close()

	client := NewClient(conn)
	go client.WriteLoop()

	decoder := json.NewDecoder(conn)

	var msg Message
	if err := decoder.Decode(&msg); err != nil {
		return
	}

	if msg.Type != "auth" {
		client.Send(Message{
			Type:  "error",
			Error: "authentication required",
		})
		return
	}

	claims, err := utils.ParseToken(msg.Token)
	if err != nil {
		client.Send(Message{
			Type:  "error",
			Error: "invalid token",
		})
		return
	}

	client.UserID = claims.UserID
	client.SessionID = "sess_" + uuid.NewString()

	hub.Register(client.UserID, client)

	log.Println("[TCP] client registered:", client.UserID, client.SessionID)

	client.Send(Message{
		Type:      "auth_ok",
		UserID:    client.UserID,
		SessionID: client.SessionID,
	})

	// giữ connection
	for {
		if err := decoder.Decode(&msg); err != nil {
			break
		}
	}

	hub.Unregister(client.UserID, client)
	log.Println("[TCP] client disconnected:", client.UserID, client.SessionID)
}
