package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLerp(t *testing.T) {
	require.EqualValues(t, 1, Lerp(1, 0, 2, -2, 4))
}
