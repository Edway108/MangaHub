// The function `GetUserProgress` establishes a gRPC connection to a server and retrieves the progress of a user identified by their ID.
package grpcclient

import (
	"context"
	"time"

	pb "MangaHub/internal/grpc/pb"

	"google.golang.org/grpc"
)

	func GetUserProgress(userID string) (*pb.UserProgressResponse, error) {
		conn, err := grpc.Dial(
			"localhost:9092",
			grpc.WithInsecure(),
		)
		if err != nil {
			return nil, err
		}
		defer conn.Close()

		client := pb.NewMangaServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		return client.GetUserProgress(ctx, &pb.UserRequest{
			UserId: userID,
		})
	}
