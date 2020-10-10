// +build test

package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
)

func TestCreateReview(t *testing.T) {
	t.Parallel()

	req := &pb.CreateReviewRequest{}

	_, err := client.CreateReview(ctx, req)

	require.NoError(t, err)
}