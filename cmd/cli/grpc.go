package main

import (
	"context"
	"fmt"
	"strconv"

	pb "MangaHub/internal/grpc/pb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "gRPC operations",
}

/*
   ======================
   gRPC GET MANGA
   ======================
*/

var grpcGetCmd = &cobra.Command{
	Use:   "get [manga_id]",
	Short: "Get manga via gRPC",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		mangaID := args[0]

		conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := pb.NewMangaServiceClient(conn)

		resp, err := client.GetManga(
			context.Background(),
			&pb.GetMangaRequest{
				MangaId: mangaID,
			},
		)
		if err != nil {
			return err
		}

		fmt.Println("ID:", resp.Id)
		fmt.Println("Title:", resp.Title)
		fmt.Println("Author:", resp.Author)
		fmt.Println("Description:", resp.Description)

		return nil
	},
}

/*
   ======================
   gRPC UPDATE PROGRESS
   ======================
*/

var grpcUpdateCmd = &cobra.Command{
	Use:   "update [user_id] [manga_id] [chapter]",
	Short: "Update manga progress via gRPC",
	Args:  cobra.ExactArgs(3),

	RunE: func(cmd *cobra.Command, args []string) error {
		userID := args[0]
		mangaID := args[1]
		chapterStr := args[2]

		chapter, err := strconv.Atoi(chapterStr)
		if err != nil {
			return fmt.Errorf("chapter must be a number")
		}

		conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := pb.NewMangaServiceClient(conn)

		resp, err := client.UpdateProgress(
			context.Background(),
			&pb.ProgressRequest{
				UserId:  userID,
				MangaId: mangaID,
				Chapter: int32(chapter),
			},
		)
		if err != nil {
			return err
		}

		if resp.Success {
			fmt.Println("Progress updated successfully via gRPC")
		} else {
			fmt.Println("Failed to update progress")
		}

		return nil
	},
}

func init() {
	grpcCmd.AddCommand(grpcGetCmd)
	grpcCmd.AddCommand(grpcUpdateCmd)
}
