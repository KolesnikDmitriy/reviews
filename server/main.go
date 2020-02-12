package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/KolesnikDmitriy/reviews/reviews"
)

const (
	port = ":50051"
)

type server struct {
	pb.ReviewsServer
}

func (s *server) CreateReview(ctx context.Context, in *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	return &pb.CreateReviewResponse{}, nil
}

func (s *server) GetRatingByItemId(ctx context.Context, in *pb.GetRatingByItemIdRequest) (*pb.GetRatingByItemIdResponse, error) {
	return &pb.GetRatingByItemIdResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReviewsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
