package perfdemo

import (
	"testing"
)

func BenchmarkFillBlock(b *testing.B) {
	if b.N < 10 {
		// we want to fill at least 10 blocks
		b.N = 10
	}
	for i := 0; i < b.N; i++ {
		// Make a block with a million tuples in it
		buf := FillBlock(1e6)
		// Tell the benchmark how many bytes that was
		b.SetBytes(int64(len(buf)))
	}
}
