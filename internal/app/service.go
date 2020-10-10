package app

import (
	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
	"google.golang.org/grpc"
)

// ReviewsService ...
type ReviewsService struct {
	pb.ReviewsServer
}

// RegisterNewReviewsService ...
func RegisterNewReviewsService(s *grpc.Server) {
	pb.RegisterReviewsServer(s, &ReviewsService{})
}
