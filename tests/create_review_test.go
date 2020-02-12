package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	pb "github.com/KolesnikDmitriy/reviews/reviews"
)

func Test_CreateReview(t *testing.T) {
	t.Parallel()

	req := &pb.CreateReviewRequest{}

	_, err := createReview(req)

	require.NoError(t, err)
}
