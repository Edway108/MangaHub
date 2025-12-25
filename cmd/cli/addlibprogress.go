package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type ProgressUpdateRequest struct {
	Chapter int `json:"chapter"`
}

var addprogressCmd = &cobra.Command{
	Use:   "addprogress",
	Short: "Manage reading progress",
}

var addprogressUpdateCmd = &cobra.Command{
	Use:   "update <manga_id> <chapter>",
	Short: "Update reading progress",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			return fmt.Errorf("usage: mangahub progress update <manga_id> <chapter>")
		}

		tokenBytes, err := os.ReadFile(".mangahub_token")
		if err != nil {
			return fmt.Errorf("not logged in (run `mangahub login` first)")
		}

		token := strings.TrimSpace(string(tokenBytes))
		if token == "" {
			return fmt.Errorf("empty token")
		}

		mangaID := args[0]
		chapter := args[1]

		payload := ProgressUpdateRequest{}
		if _, err := fmt.Sscanf(chapter, "%d", &payload.Chapter); err != nil {
			return fmt.Errorf("invalid chapter: %s", chapter)
		}

		body, _ := json.Marshal(payload)

		url := fmt.Sprintf("http://localhost:8080/progress/%s", mangaID)

		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("update progress failed: status %d", resp.StatusCode)
		}

		fmt.Printf(" Progress updated: %s => Chapter %d\n", mangaID, payload.Chapter)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addprogressCmd)
	addprogressCmd.AddCommand(addprogressUpdateCmd)
}
