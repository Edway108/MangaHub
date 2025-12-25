package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

type Manga struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Genres        string `json:"genres"`
	Status        string `json:"status"`
	TotalChapters int    `json:"total_chapters"`
	Description   string `json:"description"`
}

var mangasCmd = &cobra.Command{
	Use:   "manga",
	Short: "Manga operations",
}

var mangaSearcshCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search manga (REST API)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyword := args[0]

		token := LoadToken()
		if token == "" {
			return fmt.Errorf("not logged in, please login first")
		}

		endpoint := fmt.Sprintf(
			"http://localhost:8080/manga?search=%s",
			url.QueryEscape(keyword),
		)

		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("search failed: status %d", resp.StatusCode)
		}

		var mangas []Manga
		if err := json.NewDecoder(resp.Body).Decode(&mangas); err != nil {
			return err
		}

		if len(mangas) == 0 {
			fmt.Println("No manga found.")
			return nil
		}

		for _, m := range mangas {
			fmt.Printf(
				"- %s (%s) | %s | %s | %d chapters\n",
				m.Title,
				m.ID,
				m.Author,
				m.Genres,
				m.TotalChapters,
			)
		}

		return nil
	},
}

func init() {
	mangasCmd.AddCommand(mangaSearcshCmd)
	rootCmd.AddCommand(mangasCmd)
}
