package main

import (
	"context"
	"fmt"

	pb "MangaHub/internal/grpc/pb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

	var mangaCmd = &cobra.Command{
		Use:   "manga",
		Short: "Manga operations",
	}

	var mangaSearchCmd = &cobra.Command{
		Use:   "search [keyword]",
		Short: "Search manga",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
			if err != nil {
				return err
			}
			defer conn.Close()

			client := pb.NewMangaServiceClient(conn)

			resp, err := client.SearchManga(
				context.Background(),
				&pb.SearchRequest{
					Keyword: args[0],
					Limit:   10,
					Offset:  0,
				},
			)
			if err != nil {
				return err
			}

			for _, m := range resp.Results {
				fmt.Printf("- %s (%s)\n", m.Title, m.Id)
			}
			return nil
		},
	}

	func init() {
		mangaCmd.AddCommand(mangaSearchCmd)
	}
