package goid

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestULID(t *testing.T) {
	id1, err := NewULID().String()
	require.NoError(t, err)
	require.NotEmpty(t, id1)

	id2, err := NewULID().String()
	require.NoError(t, err)
	require.NotEmpty(t, id2)

	require.NotEqual(t, id1, id2)

	now := time.Now()
	// same user
	buf1 := bytes.NewBufferString("userid-01F1MSHQ4085724EKGNT3JKVDF")
	id3, err := NewFromTimeAndReader(now, buf1).String()
	require.NoError(t, err)
	require.NotEmpty(t, id3)

	buf2 := bytes.NewBufferString("userid-01F1MSHQ4085724EKGNT3JKVDF")
	id4, err := NewFromTimeAndReader(now, buf2).String()
	require.NoError(t, err)
	require.NotEmpty(t, id4)

	require.Equal(t, id3, id4)

	// multi user
	buf3 := bytes.NewBufferString("userid1-01F1MSHQ4085724EKGNT3JKVDF")
	id5, err := NewFromTimeAndReader(now, buf3).String()
	require.NoError(t, err)
	require.NotEmpty(t, id5)

	buf4 := bytes.NewBufferString("userid2-01F1MSHQ4085724EKGNT3JKVDF")
	id6, err := NewFromTimeAndReader(now, buf4).String()
	require.NoError(t, err)
	require.NotEmpty(t, id6)

	require.NotEqual(t, id5, id6)
}
