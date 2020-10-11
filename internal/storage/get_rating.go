package storage

import (
	"context"
	"time"
)

// GetRatingRequest ...
type GetRatingRequest struct {
	ItemID int64
}

// GetRatingResponse ...
type GetRatingResponse struct {
	Rating float32
}

// GetRating ...
func (s *Storage) GetRating(ctx context.Context, req *GetRatingRequest) (*GetRatingResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	var total, count int32
	for _, r := range s.reviews {
		if r.itemID == req.ItemID {
			total += r.score
			count++
		}
	}
	return &GetRatingResponse{Rating: float32(total) / float32(count)}, nil
}
