package tcp

import "net"

type Hub struct {
	clients map[string][]net.Conn
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string][]net.Conn),
	}

}

func (h *Hub) Register(userID string, conn net.Conn) {
	h.clients[userID] = append(h.clients[userID], conn)

}
func (h *Hub) Broadcast(userID string, msg []byte) {
	for _, conn := range h.clients[userID] {
		conn.Write(msg)
	}
}
