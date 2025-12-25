package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID   string
	Username string
	Room     string
	Conn     *websocket.Conn
	Send     chan Message
}

func (c *Client) ReadPump(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}

		msg := Message{
			Type:    "message",
			Content: string(data),
			From:    c.Username,
			Room:    c.Room,
			Client:  c,
		}

		hub.Broadcast <- msg
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		c.Conn.WriteJSON(msg)
	}
}
