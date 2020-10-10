package test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
)

const (
	address = "localhost:50051"
	timeout = time.Second
)

var (
	client pb.ReviewsClient
	ctx    context.Context
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	client = pb.NewReviewsClient(conn)

	m.Run()
}
