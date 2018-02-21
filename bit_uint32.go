package bitutil

func BigEndianToUint32WithPos(b []byte, startBit uint, bitLength uint) uint32 {
	result := uint32(0)
	startByte := startBit / 8
	rightBit := startBit + bitLength
	for i, t := range b[startByte:] {
		if i == 0 && startBit%8 > 0 {
			t = t & (1<<(8-startBit%8) - 1)
		}
		result |= uint32(t)
		if (int(startByte)+i+1)*8 >= int(rightBit) {
			break
		}
		result <<= 8
	}
	if rightBit%8 != 0 {
		result >>= (8 - (rightBit)%8)
	}
	return result
}

func Uint32ToBigEndianWithPos(b []byte, v uint32, startBit uint, bitLength uint) {
	v <<= (uint(len(b)*8) - (startBit + bitLength))
	v &= 1<<(uint(len(b)*8)-startBit) - 1
	bb := make([]byte, len(b))
	Uint32ToBigEndian(bb, v)
	for i := range b {
		b[i] |= bb[i]
	}
}

func BigEndianToUint32(b []byte) uint32 {
	result := uint32(0)
	for i := range b {
		result |= uint32(b[i]) << (8 * uint(len(b)-1-i))
	}
	return result
}

func Uint32ToBigEndian(b []byte, v uint32) {
	for i := range b {
		b[i] = byte(v >> (uint(len(b)-1-i) * 8))
	}
}

func LittleEndianToUint32(b []byte) uint32 {
	result := uint32(0)
	for i := range b {
		result |= uint32(b[i]) << (8 * uint(i))
	}
	return result
}

func Uint32ToLittleEndian(b []byte, v uint32) {
	for i := range b {
		b[i] = byte(v >> (uint(i) * 8))
	}
}
