package udp

import (
	"encoding/json"
	"net"
)

func BroadcastNewChapter(mangaID string, chapter int) {
	conn, _ := net.Dial("udp", "localhost:9100")
	defer conn.Close()
	msg := Message{
		Type:    "broadcast",
		MangaID: mangaID,
		Chapter: chapter,
	}
	data, _ := json.Marshal(msg)
	conn.Write(data)
}
