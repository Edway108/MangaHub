package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func LoadToken() string {
	data, err := os.ReadFile(".mangahub_token")
	if err != nil {
		fmt.Println(" Cannot read .mangahub_token. Please login first.")
		os.Exit(1)
	}

	token := strings.TrimSpace(string(data))
	if token == "" {
		fmt.Println(" Token file is empty. Please login again.")
		os.Exit(1)
	}

	return token
} // root: mangahub sync
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Real-time progress synchronization",
}

func init() {
	// mangahub sync connect
	syncCmd.AddCommand(&cobra.Command{
		Use:   "connect",
		Short: "Connect to TCP sync server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return SyncConnect(LoadToken())
		},
	})

	// mangahub sync monitor
	syncCmd.AddCommand(&cobra.Command{
		Use:   "monitor",
		Short: "Monitor real-time sync updates",
		RunE: func(cmd *cobra.Command, args []string) error {
			return SyncMonitor(LoadToken())
		},
	})
}
