package storage

import (
	"context"
	"sync"
)

type review struct {
	id      int64
	userID  int64
	itemID  int64
	score   int32
	message string
}

type storage interface {
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error)
	GetRating(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error)
}

// Storage ...
type Storage struct {
	storage

	reviews []review

	mu sync.Mutex
}

// NewStorage ...
func NewStorage() *Storage {
	return &Storage{
		reviews: []review{review{id: 0}},
	}
}
