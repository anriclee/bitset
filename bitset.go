package bitset

import (
	"errors"
	"math/bits"
)

type BitSet []byte

// NewBitSet ... init a bitset with a byte array whose num of bit is specified by param:size
func NewBitSet(size int64) BitSet {
	size = (size + 7) &^ 7
	return make(BitSet, size>>3)
}

func (bs BitSet) String() string {
	return string(bs)
}

func (bs BitSet) Len() int64 {
	return int64(len(bs)) << 3
}

func (bs BitSet) SetBit(bitPos int64) error {
	if bitPos > bs.Len() {
		return errors.New("bit pos out of range")
	}
	bs[bitPos>>3] |= 1 << (bitPos & 7)
	return nil
}

// RevertByteBits ... output a new BitSet which reverts bits in every byte element
func (bs BitSet) RevertByteBits() BitSet {
	dst := make(BitSet, len(bs))
	for i, b := range bs {
		dst[i] = bits.Reverse8(b)
	}
	return dst
}

func (bs BitSet) IsBitSet(bitPos int64) bool {
	if bitPos > bs.Len() {
		return false
	}
	return bs[bitPos>>3]&(1<<(bitPos&7)) != 0
}
