package test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
	"github.com/KolesnikDmitriy/reviews/test/config"
)

const timeout = time.Second

var (
	ctx    context.Context
	client pb.ReviewsClient
)

func TestMain(m *testing.M) {
	ctx = context.Background()

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	cfg := config.ProcessConfig()

	conn, _ := grpc.Dial(cfg.BaseURL, grpc.WithInsecure())
	defer conn.Close()
	client = pb.NewReviewsClient(conn)

	m.Run()
}
