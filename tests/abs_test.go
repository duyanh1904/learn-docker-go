package abs_test

import (
	"math"
	"math/big"
	"testing"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func BenchmarkBigLen(b *testing.B) {
	nigInt := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nigInt.BitLen()
	}
}
