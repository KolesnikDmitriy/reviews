package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "github.com/KolesnikDmitriy/reviews/reviews"
)

const (
	address = "localhost:50051"
	timeout = time.Second * 10
)

var (
	client pb.ReviewsClient
	ctx    = context.Background()
)

func TestMain(m *testing.M) {
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	client = pb.NewReviewsClient(conn)

	code := m.Run()
	os.Exit(code)
}

func createReview(req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return client.CreateReview(ctx, req)
}

func getRatingByItemID(req *pb.GetRatingByItemIdRequest) (*pb.GetRatingByItemIdResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return client.GetRatingByItemId(ctx, req)
}
