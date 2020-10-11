package storage

import (
	"context"
	"fmt"
	"time"
)

// CreateReviewRequest ...
type CreateReviewRequest struct {
	ItemID  int64
	UserID  int64
	Score   int32
	Message string
}

// CreateReviewResponse ...
type CreateReviewResponse struct {
	ID int64
}

// CreateReview ...
func (s *Storage) CreateReview(ctx context.Context, req *CreateReviewRequest) (*CreateReviewResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, r := range s.reviews {
		if r.itemID == req.ItemID && r.userID == req.UserID {
			return nil, fmt.Errorf("conflict")
		}
	}

	id := int64(len(s.reviews))
	fmt.Println(id)
	review := review{
		id:      id,
		itemID:  req.ItemID,
		userID:  req.UserID,
		score:   req.Score,
		message: req.Message,
	}
	s.reviews = append(s.reviews, review)
	return &CreateReviewResponse{ID: id}, nil
}
