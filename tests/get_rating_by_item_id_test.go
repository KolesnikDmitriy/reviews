package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	pb "github.com/KolesnikDmitriy/reviews/reviews"
)

func Test_GetRatingByItemId(t *testing.T) {
	t.Parallel()

	req := &pb.GetRatingByItemIdRequest{}

	_, err := getRatingByItemID(req)

	require.NoError(t, err)
}
