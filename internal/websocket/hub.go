package websocket

import (
	"strings"
	"time"
)

type Hub struct {
	Clients map[*Client]bool
	Rooms   map[string]map[*Client]bool
	History map[string][]Message

	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Rooms:      make(map[string]map[*Client]bool),
		History:    make(map[string][]Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			h.handleRegister(c)

		case c := <-h.Unregister:
			h.handleUnregister(c)

		case msg := <-h.Broadcast:
			h.handleMessage(msg)
		}
	}
}


func (h *Hub) handleRegister(c *Client) {
	h.Clients[c] = true

	if c.Room == "" {
		c.Room = "general"
	}

	if h.Rooms[c.Room] == nil {
		h.Rooms[c.Room] = make(map[*Client]bool)
	}
	h.Rooms[c.Room][c] = true

	c.Send <- Message{
		Type:    "system",
		Content: " Connected to #" + c.Room,
	}
}


func (h *Hub) handleUnregister(c *Client) {
	if _, ok := h.Clients[c]; ok {
		delete(h.Clients, c)
		delete(h.Rooms[c.Room], c)
		close(c.Send)
	}
}


func (h *Hub) handleMessage(msg Message) {
	switch msg.Type {

	case "message":
		h.broadcastToRoom(msg)

	case "command":
		h.handleCommand(msg)
	}
}


func (h *Hub) broadcastToRoom(msg Message) {

	msg.Timestamp = time.Now().Unix()
	h.History[msg.Room] = append(h.History[msg.Room], msg)

	for c := range h.Rooms[msg.Room] {
		c.Send <- msg
	}
}


func (h *Hub) handleCommand(msg Message) {
	switch msg.Command {

	case "users":
		h.sendUserList(msg)

	case "pm":
		h.handlePM(msg)

	case "manga":
		h.switchRoom(msg)

	case "history":
		h.sendHistory(msg)

	case "quit":
		h.Unregister <- msg.Client
	}
}

func (h *Hub) sendUserList(msg Message) {
	var users []string
	for c := range h.Rooms[msg.Room] {
		users = append(users, c.Username)
	}

	msg.Client.Send <- Message{
		Type:    "system",
		Content: "Online users: " + strings.Join(users, ", "),
	}
}

func (h *Hub) handlePM(msg Message) {
	for c := range h.Clients {
		if c.UserID == msg.To {
			c.Send <- Message{
				Type: "message",
				From: msg.Client.Username,
				Content:   "[PM] " + msg.Content,
				Timestamp: time.Now().Unix(),
			}
			return
		}
	}

	msg.Client.Send <- Message{
		Type:    "system",
		Content: "User not found",
	}
}
func (h *Hub) register(c *Client) {
	h.Register <- c
}

func (h *Hub) switchRoom(msg Message) {
	c := msg.Client

	delete(h.Rooms[c.Room], c)

	c.Room = msg.Room
	if h.Rooms[c.Room] == nil {
		h.Rooms[c.Room] = make(map[*Client]bool)
	}
	h.Rooms[c.Room][c] = true

	c.Send <- Message{
		Type:    "system",
		Content: "Switched to #" + c.Room,
	}
}

func (h *Hub) sendHistory(msg Message) {
	hist := h.History[msg.Room]
	if len(hist) > 10 {
		hist = hist[len(hist)-10:]
	}

	for _, m := range hist {
		msg.Client.Send <- m
	}
}
