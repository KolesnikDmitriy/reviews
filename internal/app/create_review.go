package app

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/KolesnikDmitriy/reviews/internal/storage"
	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
)

// CreateReview ...
func (s *ReviewsService) CreateReview(ctx context.Context, in *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	if in.GetItemId() <= 0 || in.GetUserId() <= 0 || in.GetScore() <= 0 || in.GetScore() > 5 {
		log.Printf("in: %v\n", in)
		return nil, grpc.Errorf(codes.InvalidArgument, "item id and user id are required fields, score should be in range 1..5\n")
	}
	res, err := s.storage.CreateReview(ctx, &storage.CreateReviewRequest{
		UserID:  in.GetUserId(),
		ItemID:  in.GetItemId(),
		Score:   in.GetScore(),
		Message: in.GetMessage(),
	})
	if err != nil {
		log.Printf("err: %v\nin: %v\n", err.Error(), in)
		code := codes.Internal
		if err.Error() == "conflict" {
			code = codes.AlreadyExists
		}
		return nil, grpc.Errorf(code, "failed to CreateReview")
	}

	return &pb.CreateReviewResponse{Id: res.ID}, nil
}
