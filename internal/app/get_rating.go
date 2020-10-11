package app

import (
	"context"

	"github.com/KolesnikDmitriy/reviews/internal/storage"
	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
)

// GetRating ...
func (s *ReviewsService) GetRating(ctx context.Context, in *pb.GetRatingRequest) (*pb.GetRatingResponse, error) {
	res, err := s.storage.GetRating(ctx, &storage.GetRatingRequest{ItemID: in.GetItemId()})
	if err != nil {
		return nil, err
	}
	return &pb.GetRatingResponse{Rating: res.Rating}, nil
}
