package fx

import (
	"slices"
	"testing"
)

//go:noinline
func use(v int) {}

func BenchmarkSliceRange(b *testing.B) {
	s := slices.Collect(Range(0, 100000))
	b.ResetTimer()

	for b.Loop() {
		sum := 0
		for _, v := range s {
			sum += v
		}
		use(sum)
	}
}

func BenchmarkSliceIter(b *testing.B) {
	s := slices.Collect(Range(0, 100000))
	b.ResetTimer()

	for b.Loop() {
		sum := 0
		for v := range slices.Values(s) {
			sum += v
		}
		use(sum)
	}
}

func BenchmarkSliceIterCallback(b *testing.B) {
	s := slices.Collect(Range(0, 100000))
	b.ResetTimer()

	for b.Loop() {
		sum := 0
		slices.Values(s)(func(v int) bool {
			sum += v
			return true
		})
		use(sum)
	}
}
