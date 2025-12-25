package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var libraryCmd = &cobra.Command{
	Use:   "library",
	Short: "Manage manga library",
}

var libraryAddCmd = &cobra.Command{
	Use:   "add <manga_id>",
	Short: "Add manga to library",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return fmt.Errorf("usage: mangahub library add <manga_id>")
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
		url := fmt.Sprintf("http://localhost:8080/library/%s", mangaID)

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			return fmt.Errorf("add library failed: status %d", resp.StatusCode)
		}

		fmt.Println(" Added to library:", mangaID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(libraryCmd)
	libraryCmd.AddCommand(libraryAddCmd)
}
