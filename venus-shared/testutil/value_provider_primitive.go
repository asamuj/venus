package testutil

import (
	crand "crypto/rand"
	"encoding/hex"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	MustRegisterDefaultValueProvier(IntProvider)

	MustRegisterDefaultValueProvier(BytesFixedProvider(defaultBytesFixedSize))
	MustRegisterDefaultValueProvier(StringInnerFixedProvider(defaultBytesFixedSize))
}

const (
	defaultBytesFixedSize = 16
)

func IntProvider(t *testing.T) int { return rand.Int() }

func IntRangedProvider(min, max int) func(*testing.T) int {
	return func(t *testing.T) int {
		gap := max - min
		if gap <= 0 {
			t.Fatalf("invalid range [%d, %d)", min, max)
		}

		return min + rand.Intn(gap)
	}
}

func BytesFixedProvider(size int) func(*testing.T) []byte {
	return func(t *testing.T) []byte {
		b := make([]byte, size)
		_, err := crand.Read(b[:])
		require.NoError(t, err)
		return b
	}
}

func BytesAtMostProvider(size int) func(*testing.T) []byte {
	return func(t *testing.T) []byte {
		b := make([]byte, rand.Intn(size))
		_, err := crand.Read(b[:])
		require.NoError(t, err)
		return b
	}
}

func StringInnerFixedProvider(size int) func(*testing.T) string {
	return func(t *testing.T) string {
		return hex.EncodeToString(BytesFixedProvider(size)(t))
	}
}

func StringInnerAtMostProvider(size int) func(*testing.T) string {
	return func(t *testing.T) string {
		return hex.EncodeToString(BytesFixedProvider(size)(t))
	}
}
