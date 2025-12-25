package main

import (
	"context"
	"fmt"

	pb "MangaHub/internal/grpc/pb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var mangaID string
var chapter int32

var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Progress tracking",
}

var progressUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update reading progress",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := pb.NewMangaServiceClient(conn)

		userID, err := getUserIDFromToken()
		if err != nil {
			return err
		}

		_, err = client.UpdateProgress(
			context.Background(),
			&pb.ProgressRequest{
				UserId:  userID,
				MangaId: mangaID,
				Chapter: chapter,
			},
		)

		if err != nil {
			return err
		}

		fmt.Println(" Progress updated successfully")
		return nil
	},
}

func init() {
	progressUpdateCmd.Flags().StringVar(&mangaID, "manga-id", "", "Manga ID")
	progressUpdateCmd.Flags().Int32Var(&chapter, "chapter", 0, "Chapter number")
	progressUpdateCmd.MarkFlagRequired("manga-id")
	progressUpdateCmd.MarkFlagRequired("chapter")

	progressCmd.AddCommand(progressUpdateCmd)
	grpcCmd.AddCommand(progressCmd)
}
