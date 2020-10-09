package app

import (
	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
	"google.golang.org/grpc"
)

// ReviewsService ...
type ReviewsService struct {
	pb.ReviewsServer
}

// NewReviewsService ...
func NewReviewsService() *ReviewsService {
	return &ReviewsService{}
}

// RegisterReviewsServer ...
func (r *ReviewsService) RegisterReviewsServer(s *grpc.Server) {
	pb.RegisterReviewsServer(s, r)
}
