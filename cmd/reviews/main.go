package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/DmitriiKolesnik/reviews/internal/app"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service := app.NewReviewsService()
	service.RegisterReviewsServer(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
