// +build test

package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	pb "github.com/KolesnikDmitriy/reviews/pkg/api"
)

func TestGetRating(t *testing.T) {
	t.Parallel()

	req := &pb.GetRatingRequest{}

	_, err := client.GetRating(ctx, req)

	require.NoError(t, err)
}
