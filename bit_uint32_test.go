package bitutil

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBigEndianToUint32WithPos(t *testing.T) {
	base4 := uint32(0xc2<<24 | 0x13<<16 | 0x14<<8 | 0x15)
	base3 := uint32(0xc2<<16 | 0x13<<8 | 0x14)
	testCases := []struct {
		b         []byte
		startBit  uint
		bitLength uint
		result    uint32
	}{
		{
			b:         []byte{0xc2, 0x13, 0x14, 0x15},
			startBit:  0,
			bitLength: 32,
			result:    base4,
		},
		{
			b:         []byte{0xc2, 0x13, 0x14, 0x15},
			startBit:  0,
			bitLength: 31,
			result:    base4 >> 1,
		},
		{
			b:         []byte{0xc2, 0x13, 0x14, 0x15},
			startBit:  1,
			bitLength: 31,
			result:    base4 & (1<<31 - 1),
		},
		{
			b:         []byte{0xc2, 0x13, 0x14, 0x15},
			startBit:  1,
			bitLength: 30,
			result:    base4 & (1<<31 - 1) >> 1,
		},
		{
			b:         []byte{0xc2, 0x13, 0x14, 0x15},
			startBit:  9,
			bitLength: 14,
			result:    base4 & (1<<23 - 1) >> 9,
		},
		{
			b:         []byte{0xc2, 0x13, 0x14},
			startBit:  0,
			bitLength: 24,
			result:    base3,
		},
		{
			b:         []byte{0xc2, 0x13, 0x14},
			startBit:  0,
			bitLength: 23,
			result:    base3 >> 1,
		},
		{
			b:         []byte{0xc2, 0x13, 0x14},
			startBit:  1,
			bitLength: 23,
			result:    base3 & (1<<23 - 1),
		},
		{
			b:         []byte{0xc2, 0x13, 0x14},
			startBit:  1,
			bitLength: 22,
			result:    base3 & (1<<23 - 1) >> 1,
		},
		{
			b:         []byte{0xc2},
			startBit:  1,
			bitLength: 6,
			result:    0x42 >> 1,
		},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			r := BigEndianToUint32WithPos(tt.b, tt.startBit, tt.bitLength)
			if r != tt.result {
				t.Errorf("want %x but %x", tt.result, r)
			}
		})
	}
}

func TestUint32ToBigEndianWithPos(t *testing.T) {
	testCases := []struct {
		b         []byte
		v         uint32
		startBit  uint
		bitLength uint
		result    []byte
	}{
		{
			b:         make([]byte, 4),
			v:         0xc2<<24 | 0x13<<16 | 0x14<<8 | 0x15,
			startBit:  0,
			bitLength: 32,
			result:    []byte{0xc2, 0x13, 0x14, 0x15},
		},
		{
			b:         make([]byte, 4),
			v:         (0xc2<<24 | 0x13<<16 | 0x14<<8 | 0x15) >> 1,
			startBit:  0,
			bitLength: 31,
			result:    []byte{0xc2, 0x13, 0x14, 0x15 | 1 ^ 1},
		},
		{
			b:         make([]byte, 4),
			v:         0x42<<24 | 0x13<<16 | 0x14<<8 | 0x15,
			startBit:  1,
			bitLength: 31,
			result:    []byte{0x42, 0x13, 0x14, 0x15},
		},
		{
			b:         make([]byte, 4),
			v:         (0x42<<24 | 0x13<<16 | 0x14<<8 | 0x15) >> 1,
			startBit:  1,
			bitLength: 30,
			result:    []byte{0x42, 0x13, 0x14, 0x15 | 1 ^ 1},
		},
		{
			b:         []byte{1, 0, 0, 0},
			v:         (0x42<<8 | 0x13) >> 1,
			startBit:  9,
			bitLength: 14,
			result:    []byte{1, 0x42, 0x13 | 1 ^ 1, 0},
		},
		{
			b:         make([]byte, 3),
			v:         0xc2<<16 | 0x13<<8 | 0x14,
			startBit:  0,
			bitLength: 24,
			result:    []byte{0xc2, 0x13, 0x14},
		},
		{
			b:         make([]byte, 3),
			v:         (0xc2<<16 | 0x13<<8 | 0x14) >> 1,
			startBit:  0,
			bitLength: 23,
			result:    []byte{0xc2, 0x13, 0x14 | 1 ^ 1},
		},
		{
			b:         make([]byte, 3),
			v:         0x42<<16 | 0x13<<8 | 0x14,
			startBit:  1,
			bitLength: 23,
			result:    []byte{0x42, 0x13, 0x14},
		},
		{
			b:         make([]byte, 3),
			v:         (0x42<<16 | 0x13<<8 | 0x14) >> 1,
			startBit:  1,
			bitLength: 22,
			result:    []byte{0x42, 0x13, 0x14 | 1 ^ 1},
		},
		{
			b:         make([]byte, 1),
			v:         0x42 >> 1,
			startBit:  1,
			bitLength: 6,
			result:    []byte{0x42 | 1 ^ 1},
		},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			Uint32ToBigEndianWithPos(tt.b, tt.v, tt.startBit, tt.bitLength)
			if !bytes.Equal(tt.b, tt.result) {
				t.Errorf("want %x but %x", tt.result, tt.b)
			}
		})
	}
}
