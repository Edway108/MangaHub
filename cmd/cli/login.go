// The function `loginCmd` in the Go code snippet is used to authenticate a user and store the JWT token received in a file.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type LoginResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

var loginCmd = &cobra.Command{
	Use:   "login <username> <password>",
	Short: "Login to MangaHub",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			return fmt.Errorf("usage: app login <username> <password>")
		}

		payload := map[string]string{
			"username": args[0],
			"password": args[1],
		}

		body, _ := json.Marshal(payload)

		resp, err := http.Post(
			"http://localhost:8080/auth/login",
			"application/json",
			bytes.NewBuffer(body),
		)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("login failed: status %d", resp.StatusCode)
		}

		var result struct {
			Token string `json:"token"`
		}

		json.NewDecoder(resp.Body).Decode(&result)

		if result.Token == "" {
			return fmt.Errorf("empty token received")
		}

		fmt.Println("JWT:", result.Token)
		os.WriteFile(".mangahub_token", []byte(result.Token), 0644)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
