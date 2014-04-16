package perfdemo

import (
	"bytes"
)

// FillBlock returns a block of populated tuples
func FillBlock(count int) []byte {
	// buf is out result buffer
	var buf []byte

	// append 'count' number of tuples to the buffer
	for i := 0; i < count; i++ {
		buf = append(buf, tupleBytes(i)...)
	}

	return buf
}

// tupleBytes generates an encoded tuple with the format:
// CaptureTime uint64 (8 bytes)
// PutTime     uint64 (8 bytes)
// Encoded     []byte (50 bytes)
func tupleBytes(i int) []byte {
	captureTime := uint64(i)
	putTime := uint64(i)
	tuple := make([]byte, 8+8)
	PutLittleEndianUint64(tuple, 0, captureTime)
	PutLittleEndianUint64(tuple, 8, putTime)
	encoded := bytes.Repeat([]byte{1}, 50)
	tuple = append(tuple, encoded...)
	return tuple
}

func PutLittleEndianUint64(b []byte, offset int, v uint64) {
	b[offset] = byte(v)
	b[offset+1] = byte(v >> 8)
	b[offset+2] = byte(v >> 16)
	b[offset+3] = byte(v >> 24)
	b[offset+4] = byte(v >> 32)
	b[offset+5] = byte(v >> 40)
	b[offset+6] = byte(v >> 48)
	b[offset+7] = byte(v >> 56)
}
