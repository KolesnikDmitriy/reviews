package app

import (
	"context"

	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
)

// CreateReview ...
func (s *ReviewsService) CreateReview(ctx context.Context, in *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	return &pb.CreateReviewResponse{}, nil
}
