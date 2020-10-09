package app

import (
	"context"

	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
)

// GetRating ...
func (s *ReviewsService) GetRating(ctx context.Context, in *pb.GetRatingRequest) (*pb.GetRatingResponse, error) {
	return &pb.GetRatingResponse{}, nil
}
