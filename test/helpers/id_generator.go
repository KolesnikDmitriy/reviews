package helpers

import (
	"math/rand"
	"testing"
)

// NewID generate random number and log it
func NewID(t *testing.T) int64 {
	t.Helper()
	id := rand.Int63()
	t.Logf("newID: %v\n", id)
	return id
}
