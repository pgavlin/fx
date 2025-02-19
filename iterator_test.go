package fx

import (
	"slices"
	"testing"
)

func BenchmarkSliceRange(b *testing.B) {
	s := slices.Collect(Range(0, 100000))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sum := 0
		for _, v := range s {
			sum += v
		}
	}
}

func BenchmarkSliceIter(b *testing.B) {
	s := slices.Collect(Range(0, 100000))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sum := 0
		for v := range slices.Values(s) {
			sum += v
		}
	}
}
