package goid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUUID(t *testing.T) {
	str, err := NewUUID().String()
	require.NoError(t, err)
	require.Equal(t, 36, len(str))

	str2, err := NewUUID().String()
	require.NoError(t, err)
	require.NotEqual(t, str, str2)

	str, err = NewUUID().Base58Encode()
	require.NoError(t, err)
	require.True(t, len(str) <= 22)

	require.True(t, len(EncodeBase58(str2)) <= 22)
}
