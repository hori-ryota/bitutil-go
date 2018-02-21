package bitutil

func ByteToUint8WithPos(b byte, startBit uint, bitLength uint) uint8 {
	if bitLength == 8 {
		return b
	}
	return b & uint8(1<<(8-startBit)-1) >> (8 - uint8(startBit+bitLength))
}

func Uint8ToByteWithPos(v uint8, startBit uint, bitLength uint) byte {
	if bitLength == 8 {
		return v
	}
	return v << (8 - uint8(startBit+bitLength)) & uint8(1<<(8-startBit)-1)
}
