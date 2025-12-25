package udp

type Message struct {
	Type    string `json:"type"` // "notify"
	UserID  string `json:"user_id"`
	MangaID string `json:"manga_id,omitempty"`
	Content string `json:"content"` // message
}
type Notification struct {
	Type      string `json:"type"` // "register", "notify"
	MangaID   string `json:"manga_id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
