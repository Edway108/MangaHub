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
		var incoming Message
		err := c.Conn.ReadJSON(&incoming)
		if err != nil {
			return
		}

		incoming.Type = "message"
		incoming.User = c.Username
		incoming.From = c.Username
		incoming.Room = c.Room
		incoming.Client = c

		hub.Broadcast <- incoming
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		c.Conn.WriteJSON(msg)
	}
}
