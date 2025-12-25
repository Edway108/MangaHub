package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func SyncMonitor(token string) error {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		return err
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	encoder.Encode(map[string]string{
		"type":  "auth",
		"token": token,
	})

	var localSession string

	fmt.Println("Monitoring real-time sync updates...")

	for {
		var msg map[string]interface{}
		if err := decoder.Decode(&msg); err != nil {
			return err
		}

		switch msg["type"] {
		case "auth_ok":
			localSession = msg["session_id"].(string)
			fmt.Println("Connected. Session:", localSession)

		case "progress_update":
			if msg["session_id"] == localSession {
				continue
			}

			fmt.Printf(
				"[%s] Sync: %s → Chapter %v\n",
				time.Now().Format("15:04:05"),
				msg["manga_id"],
				msg["chapter"],
			)
		}
	}
}
