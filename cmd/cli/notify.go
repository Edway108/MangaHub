package main

import (
	"encoding/json"
	"fmt"
	"net"

	"MangaHub/internal/udp"

	"github.com/spf13/cobra"
)

var notifySubscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to UDP notifications",
	RunE: func(cmd *cobra.Command, args []string) error {

		addr, _ := net.ResolveUDPAddr("udp", "localhost:9091")
		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			return err
		}

		// REGISTER
		reg := udp.Notification{Type: "register"}
		data, _ := json.Marshal(reg)
		conn.Write(data)

		fmt.Println("Subscribed. Listening... (Ctrl+C to exit)")

		// LISTEN
		buf := make([]byte, 2048)
		for {
			n, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				return err
			}

			var msg udp.Notification
			_ = json.Unmarshal(buf[:n], &msg)

			fmt.Printf(" %s\n", msg.Message)
		}
	},
}

func init() {
	rootCmd.AddCommand(notifySubscribeCmd)
}
