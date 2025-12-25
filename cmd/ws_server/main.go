package main

import (
	"log"
	"net/http"

	ws "MangaHub/internal/websocket"
	"MangaHub/pkg/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	hub := ws.NewHub()
	go hub.Run()

	http.HandleFunc("/ws/chat", func(w http.ResponseWriter, r *http.Request) {

		tokenStr := r.URL.Query().Get("token")
		if tokenStr == "" {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WS upgrade error:", err)
			return
		}

		client := &ws.Client{
			UserID:   claims.UserID,
			Username: claims.Username,
			Room:     "general",
			Conn:     conn,
			Send:     make(chan ws.Message, 256),
		}

		hub.Register <- client

		log.Printf("User %s (%s) connected\n", client.Username, client.UserID)

		go client.WritePump()
		go client.ReadPump(hub)
	})

	log.Println("WebSocket chat running on :9093")
	log.Fatal(http.ListenAndServe(":9093", nil))
}
