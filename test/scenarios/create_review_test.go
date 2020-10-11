// +build test

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
	"github.com/KolesnikDmitriy/reviews/test/helpers"
)

func TestCreateDefaultReview(t *testing.T) {
	t.Parallel()

	req := &pb.CreateReviewRequest{
		UserId:  helpers.NewID(t),
		ItemId:  helpers.NewID(t),
		Score:   5,
		Message: "good price",
	}

	res, err := client.CreateReview(ctx, req)

	require.NoError(t, err)
	assert.NotZero(t, res.GetId(), "Id")
}

func TestCreateReviewWithoutMessage(t *testing.T) {
	t.Parallel()

	req := &pb.CreateReviewRequest{
		UserId: helpers.NewID(t),
		ItemId: helpers.NewID(t),
		Score:  1,
	}

	res, err := client.CreateReview(ctx, req)

	require.NoError(t, err)
	assert.NotZero(t, res.GetId(), "Id")
}

func TestDuplicateReview(t *testing.T) {
	t.Parallel()

	req := &pb.CreateReviewRequest{
		UserId: helpers.NewID(t),
		ItemId: helpers.NewID(t),
		Score:  5,
	}
	createReview(t, req)

	_, err := client.CreateReview(ctx, req)

	require.Error(t, err)
	assert.Equal(t, grpc.Code(err), codes.AlreadyExists)
	assert.Contains(t, err.Error(), "failed to CreateReview")
}

func TestValidateionFileds(t *testing.T) {
	testCases := []struct {
		name string
		req  pb.CreateReviewRequest
	}{
		{name: "empty request", req: pb.CreateReviewRequest{}},
		{name: "empty item id", req: pb.CreateReviewRequest{
			UserId: helpers.NewID(t),
			Score:  1,
		}},
		{name: "empty user id", req: pb.CreateReviewRequest{
			ItemId: helpers.NewID(t),
			Score:  5,
		}},
		{name: "empty score", req: pb.CreateReviewRequest{
			UserId: helpers.NewID(t),
			ItemId: helpers.NewID(t),
		}},
		{name: "negative score", req: pb.CreateReviewRequest{
			UserId: helpers.NewID(t),
			ItemId: helpers.NewID(t),
			Score:  -1,
		}},
		{name: "negative item id", req: pb.CreateReviewRequest{
			UserId: helpers.NewID(t),
			ItemId: -1,
			Score:  4,
		}},
		{name: "negative score", req: pb.CreateReviewRequest{
			UserId: -1,
			ItemId: helpers.NewID(t),
			Score:  3,
		}},
		{name: "high score", req: pb.CreateReviewRequest{
			UserId: helpers.NewID(t),
			ItemId: helpers.NewID(t),
			Score:  6,
		}},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := client.CreateReview(ctx, &tc.req)

			require.Error(t, err)
			assert.Equal(t, grpc.Code(err), codes.InvalidArgument, "code")
			assert.Contains(t, err.Error(), "item id and user id are required fields, score should be in range 1..5")
		})
	}
}

func createReview(t *testing.T, req *pb.CreateReviewRequest) int64 {
	t.Helper()
	res, err := client.CreateReview(ctx, req)
	require.NoError(t, err)
	return res.GetId()
}
