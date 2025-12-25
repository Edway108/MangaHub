package tcp

type Message struct {
	Type           string `json:"type"`
	UserID         string `json:"user_id,omitempty"`
	MangaID        string `json:"manga_id,omitempty"`
	CurrentChapter int    `json:"current_chapter,omitempty"`
}
