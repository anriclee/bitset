package bitset

import (
	"errors"
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
	bs[bitPos>>3] |= 1 << (7 - bitPos&7)
	return nil
}

func (bs BitSet) IsBitSet(bitPos int64) bool {
	if bitPos > bs.Len() {
		return false
	}
	return bs[bitPos>>3]&(1<<(7-bitPos&7)) != 0
}
