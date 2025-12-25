package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "MangaHub/internal/grpc/pb"
	"MangaHub/internal/udp"
	"MangaHub/pkg/database"

	"google.golang.org/grpc"
)

type MangaServer struct {
	pb.UnimplementedMangaServiceServer
	db        *sql.DB
}

func (s *MangaServer) GetUserProgress(
	ctx context.Context,
	req *pb.UserRequest,
) (*pb.UserProgressResponse, error) {

	rows, err := s.db.Query(`
		SELECT manga_id, current_chapter, status
		FROM user_progress
		WHERE user_id = ?`,
		req.UserId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*pb.UserProgress
	for rows.Next() {
		p := &pb.UserProgress{}
		if err := rows.Scan(
			&p.MangaId,
			&p.CurrentChapter,
			&p.Status,
		); err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	return &pb.UserProgressResponse{
		Progress: list,
	}, nil
}

func (s *MangaServer) GetManga(
	ctx context.Context,
	req *pb.GetMangaRequest,
) (*pb.MangaResponse, error) {

	row := s.db.QueryRow(`
		SELECT id, title, author, description
		FROM manga WHERE id = ?`,
		req.MangaId,
	)

	var m pb.MangaResponse
	if err := row.Scan(
		&m.Id, &m.Title, &m.Author, &m.Description,
	); err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *MangaServer) SearchManga(
	ctx context.Context,
	req *pb.SearchRequest,
) (*pb.SearchResponse, error) {

	rows, err := s.db.Query(`
		SELECT id, title, author, description,total_chapter
		FROM manga
		WHERE title LIKE ?
		LIMIT ? OFFSET ?`,
		"%"+req.Keyword+"%",
		req.Limit,
		req.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*pb.MangaResponse
	for rows.Next() {
		var m pb.MangaResponse
		rows.Scan(&m.Id, &m.Title, &m.Author, &m.Description)
		results = append(results, &m)
	}

	return &pb.SearchResponse{
		Results: results,
	}, nil
}

func (s *MangaServer) UpdateProgress(
	ctx context.Context,
	req *pb.ProgressRequest,
) (*pb.ProgressResponse, error) {

	_ = udp.Broadcast(
		fmt.Sprintf(
			"User %s updated %s to chapter %d",
			req.UserId, req.MangaId, req.Chapter,
		),
	)

	return &pb.ProgressResponse{Success: true}, nil
}

func main() {
	db := database.InitDB("../api_server/mangahub.db")

	lis, _ := net.Listen("tcp", ":9092")
	grpcServer := grpc.NewServer()

	pb.RegisterMangaServiceServer(
		grpcServer,
		&MangaServer{
			db:        db,},
	)

	log.Println("gRPC MangaService on :9092")
	grpcServer.Serve(lis)
}
