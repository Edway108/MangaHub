package tcp

import (
	"log"
	"sync"
)


type Hub struct {
	mu      sync.RWMutex
	clients map[string]map[*Client]bool // userID → clients
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Register(userID string, c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.clients[userID] == nil {
		h.clients[userID] = make(map[*Client]bool)
	}
	h.clients[userID][c] = true

	log.Printf("[TCP] register user=%s total_clients=%d\n",
		userID, len(h.clients[userID]))
}

func (h *Hub) Unregister(userID string, c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.clients[userID] != nil {
		delete(h.clients[userID], c)
		if len(h.clients[userID]) == 0 {
			delete(h.clients, userID)
		}
	}

	log.Printf("[TCP] unregister user=%s\n", userID)
}

func (h *Hub) Broadcast(userID string, msg Message) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	clients, ok := h.clients[userID]
	if !ok || len(clients) == 0 {
		log.Printf("[TCP] broadcast skipped: no clients for user=%s\n", userID)
		return
	}

	log.Printf("[TCP] broadcasting to %d clients for user=%s\n",
		len(clients), userID)

	for c := range clients {
		if c.SessionID == msg.SessionID {
			continue
		}
		c.Send(msg)
	}
}
