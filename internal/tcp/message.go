package tcp

type Message struct {
    Type string `json:"type"`

    // auth
    Token string `json:"token,omitempty"`

    // progress
    MangaID string `json:"manga_id,omitempty"`
    Chapter int    `json:"chapter,omitempty"`

    // meta
    UserID    string `json:"user_id,omitempty"`
    SessionID string `json:"session_id,omitempty"`
    Error     string `json:"error,omitempty"`
}
