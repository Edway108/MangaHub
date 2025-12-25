package tcp

import (
	"encoding/json"
	"log"
	"net"
)

func SendProgressSync(userID string, mangaID string, chapter int) {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("Fail to connect to the server ", err)

	}
	defer conn.Close()

	msg := Message{
		Type:           "progress_sync",
		UserID:         userID,
		MangaID:        mangaID,
		CurrentChapter: chapter,
	}

	data, _ := json.Marshal(msg)
	conn.Write(data)
}
