package tcp

import (
	"encoding/json"
	"net"
)

type Client struct {
    Conn      net.Conn
    UserID    string
    SessionID string
    SendChan  chan Message
}

func NewClient(conn net.Conn) *Client {
    return &Client{
        Conn:     conn,
        SendChan: make(chan Message, 16),
    }
}

func (c *Client) Send(msg Message) {
    select {
    case c.SendChan <- msg:
    default:
        // drop message if client lag/dead
    }
}

func (c *Client) WriteLoop() {
    encoder := json.NewEncoder(c.Conn)
    for msg := range c.SendChan {
        if err := encoder.Encode(msg); err != nil {
            return
        }
    }
}
