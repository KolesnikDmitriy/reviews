package app

import (
	"google.golang.org/grpc"

	"github.com/KolesnikDmitriy/reviews/internal/storage"
	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
)

// ReviewsService ...
type ReviewsService struct {
	pb.ReviewsServer
	storage *storage.Storage
}

// NewReviewsService ...
func NewReviewsService() *ReviewsService {
	r := ReviewsService{
		storage: storage.NewStorage(),
	}
	return &r
}

// RegisterNewReviewsService ...
func RegisterNewReviewsService(s *grpc.Server) {
	pb.RegisterReviewsServer(s, NewReviewsService())
}
