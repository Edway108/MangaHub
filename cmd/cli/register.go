package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var RegisterCmd = &cobra.Command{
	Use:   "register <username> <password>",
	Short: "register to MangaHub",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			return fmt.Errorf("usage: app register <username> <password>")
		}

		payload := map[string]string{
			"username": args[0],
			"password": args[1],
		}

		body, _ := json.Marshal(payload)

		resp, err := http.Post(
			"http://localhost:8080/auth/register",
			"application/json",
			bytes.NewBuffer(body),
		)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
			return fmt.Errorf("register failed: status %d", resp.StatusCode)
		}

		fmt.Printf("register successfully with %s ", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(RegisterCmd)
}
