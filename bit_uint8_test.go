package bitutil

import (
	"fmt"
	"testing"
)

func TestByteToUint8WithPos(t *testing.T) {
	for _, tt := range []struct {
		b         byte
		v         uint8
		startBit  uint
		bitLength uint
	}{
		{
			b:         0x3,
			v:         3,
			startBit:  6,
			bitLength: 2,
		},
		{
			b:         0x3 | 0xf0,
			v:         3,
			startBit:  6,
			bitLength: 2,
		},
		{
			b:         1<<7 | 1<<6,
			v:         3,
			startBit:  0,
			bitLength: 2,
		},
		{
			b:         0xff,
			v:         3,
			startBit:  0,
			bitLength: 2,
		},
	} {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			v := ByteToUint8WithPos(tt.b, tt.startBit, tt.bitLength)
			if v != tt.v {
				t.Errorf("want %x but %x", tt.v, v)
			}
		})
	}
}

func TestUint8ToByteWithPos(t *testing.T) {
	for _, tt := range []struct {
		b         byte
		v         uint8
		startBit  uint
		bitLength uint
	}{
		{
			b:         0x3,
			v:         3,
			startBit:  6,
			bitLength: 2,
		},
		{
			b:         1<<7 | 1<<6,
			v:         3,
			startBit:  0,
			bitLength: 2,
		},
	} {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			b := Uint8ToByteWithPos(tt.v, tt.startBit, tt.bitLength)
			if b != tt.b {
				t.Errorf("want %x but %x", tt.b, b)
			}
		})
	}
}
