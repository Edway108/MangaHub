package main

import (
	"context"
	"fmt"

	pb "MangaHub/internal/grpc/pb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "gRPC operations",
}

var grpcGetCmd = &cobra.Command{
	Use:  "get [manga_id]",
	Args: cobra.ExactArgs(1),

	Short: "Get manga via gRPC",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing manga id ")
		}

		mangaID := args[0]

		conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := pb.NewMangaServiceClient(conn)

		resp, err := client.GetManga(
			context.Background(),
			&pb.GetMangaRequest{MangaId: mangaID},
		)
		if err != nil {
			return err
		}

		fmt.Println(resp.Title, "-", resp.Author, "-", resp.Description)
		return nil
	},
}

func init() {
	grpcCmd.AddCommand(grpcGetCmd)
}
