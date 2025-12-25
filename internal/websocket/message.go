package websocket

type Message struct {
	Type      string `json:"type"`
	User      string `json:"user,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	Room      string `json:"room,omitempty"`
	Command   string `json:"command,omitempty"`
	Content   string `json:"content,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`

	// dùng nội bộ server
	Client *Client `json:"-"`
}
