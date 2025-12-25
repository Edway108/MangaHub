package models

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Pasword   string `json:"password,omitempty"`
	CreatedAt string `json:"created_at"`
}
