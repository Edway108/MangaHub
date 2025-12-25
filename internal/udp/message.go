package udp

type Message struct {
	Type     string `json:"type"`
	MangaID  string `json:"manga_id,omitempty"`
	Chapter  int    `json:"chapter,omitempty"`
}
