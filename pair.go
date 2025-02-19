package fx

import "iter"

type Pair[T, U any] struct {
	Fst T
	Snd U
}

func (p Pair[T, U]) Unpack() (T, U) {
	return p.Fst, p.Snd
}

func NewPair[T, U any](fst T, snd U) Pair[T, U] {
	return Pair[T, U]{Fst: fst, Snd: snd}
}

func Pairs[K, V any](it iter.Seq2[K, V]) iter.Seq[Pair[K, V]] {
	return func(yield func(p Pair[K, V]) bool) {
		for k, v := range it {
			if !yield(NewPair(k, v)) {
				return
			}
		}
	}
}

func Seq2[K, V any](it iter.Seq[Pair[K, V]]) iter.Seq2[K, V] {
	return func(yield func(k K, v V) bool) {
		for p := range it {
			if !yield(p.Unpack()) {
				return
			}
		}
	}
}
