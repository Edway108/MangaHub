package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

type WSMessage struct {
	Type      string `json:"type"`
	Command   string `json:"command,omitempty"`
	Room      string `json:"room,omitempty"`
	To        string `json:"to,omitempty"`
	From      string `json:"from,omitempty"`
	Content   string `json:"content,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Join MangaHub chat system",
	Run: func(cmd *cobra.Command, args []string) {
		RunChat()
	},
}

func RunChat() {
	tokenBytes, err := os.ReadFile(".mangahub_token")
	if err != nil {
		fmt.Println(" Please login first: app login")
		return
	}
	token := strings.TrimSpace(string(tokenBytes))

	room := "general"

	url := "ws://localhost:9093/ws/chat?token=" + token
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("WebSocket error:", err)
	}
	defer conn.Close()

	fmt.Println(" Connected to General Chat")
	fmt.Println("Chat Room: #general")
	fmt.Println("Type /help for commands or /quit to leave")

	go func() {
		for {
			var msg WSMessage
			if err := conn.ReadJSON(&msg); err != nil {
				fmt.Println("\n Disconnected from chat server")
				os.Exit(0)
			}

			switch msg.Type {
			case "message":
				printChatMessage(msg)

			case "system":
				fmt.Println(msg.Content)
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "/") {
			handleCommand(conn, line, &room)
			continue
		}

		// normal chat message
		conn.WriteMessage(websocket.TextMessage, []byte(line))

	}
}

func handleCommand(conn *websocket.Conn, line string, room *string) {
	parts := strings.SplitN(line, " ", 3)
	cmd := parts[0]

	switch cmd {

	case "/help":
		fmt.Println(`
Chat Commands:
/help                  Show this help
/users                 List online users
/pm <user> <message>   Private message
/manga <id>            Switch manga chat
/quit                  Leave chat
`)

	case "/users":
		conn.WriteMessage(
			websocket.TextMessage,
			[]byte("/users"),
		)

	case "/pm":
		if len(parts) < 3 {
			fmt.Println("Usage: /pm <user> <message>")
			return
		}

		conn.WriteJSON(WSMessage{
			Type:    "command",
			Command: "pm",
			To:      parts[1],
			Content: parts[2],
		})

	case "/manga":
		if len(parts) < 2 {
			fmt.Println("Usage: /manga <id>")
			return
		}

		*room = parts[1]

		conn.WriteJSON(WSMessage{
			Type:    "command",
			Command: "manga",
			Room:    *room,
		})

		fmt.Println(" Switched to " + *room)

	case "/quit":
		fmt.Println("Leaving chat...")
		conn.Close()
		os.Exit(0)

	default:
		fmt.Println("Unknown command. Type /help")
	}
}

func printChatMessage(msg WSMessage) {
	t := time.Unix(msg.Timestamp, 0).Format("15:04")
	fmt.Printf("[%s] %s: %s\n", t, msg.From, msg.Content)

}

func init() {
	rootCmd.AddCommand(chatCmd)
}
