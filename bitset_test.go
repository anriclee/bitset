package bitset

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBitSet_Len(t *testing.T) {
	testCases := []struct {
		Input    int64
		Expected int
	}{
		{
			Input:    0,
			Expected: 0,
		},
		{
			Input:    7,
			Expected: 1,
		},
		{
			Input:    9,
			Expected: 2,
		},
		{
			Input:    16,
			Expected: 2,
		},
		{
			Input:    10001,
			Expected: 1251,
		},
	}

	for _, tc := range testCases {
		set := NewBitSet(tc.Input)
		require.Equal(t, tc.Expected, len(set))
	}
}

func TestBitSet_SetBit(t *testing.T) {
	setPos := 15
	set := NewBitSet(int64(80))

	require.Equal(t, set.Len(), int64(80))
	require.Equal(t, len(set), 10)

	err := set.SetBit(int64(setPos))
	require.Nil(t, err)
	res := set.IsBitSet(int64(setPos))
	require.True(t, res)
	for i := 0; i < 80; i++ {
		if i == setPos {
			continue
		}
		res = set.IsBitSet(int64(i))
		require.False(t, res)
	}
}

func TestBitSet_Revert(t *testing.T) {
	setPos := 2
	set := NewBitSet(int64(80))

	err := set.SetBit(int64(setPos))
	require.Nil(t, err)

	reversed := set.RevertByteBits()
	require.True(t, reversed.IsBitSet(5))
}
