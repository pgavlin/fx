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

func ToSlice[T any](it iter.Seq[T]) []T {
	var s []T
	for v := range it {
		s = append(s, v)
	}
	return s
}

func TrySlice[T any](it iter.Seq[Result[T]]) ([]T, error) {
	var s []T
	for v := range it {
		v, err := v.Unpack()
		if err != nil {
			return nil, err
		}
		s = append(s, v)
	}
	return s, nil
}

func ToSet[T comparable](it iter.Seq[T]) Set[T] {
	s := Set[T]{}
	for v := range it {
		s.Add(v)
	}
	return s
}

func TrySet[T comparable](it iter.Seq[Result[T]]) (Set[T], error) {
	s := Set[T]{}
	for v := range it {
		v, err := v.Unpack()
		if err != nil {
			return Set[T]{}, err
		}
		s.Add(v)
	}
	return s, nil
}

func ToMap[K comparable, V any](it iter.Seq2[K, V]) map[K]V {
	m := map[K]V{}
	for k, v := range it {
		m[k] = v
	}
	return m
}

func TryMap[K comparable, V any](it iter.Seq2[K, Result[V]]) (map[K]V, error) {
	m := map[K]V{}
	for k, v := range it {
		v, err := v.Unpack()
		if err != nil {
			return nil, err
		}
		m[k] = v
	}
	return m, nil
}

func IterSlice[T any](ts []T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for _, t := range ts {
			if !yield(t) {
				return
			}
		}
	}
}

func IterList[T any, L List[T]](ts L) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for i := 0; i < ts.Len(); i++ {
			if !yield(ts.At(i)) {
				return
			}
		}
	}
}

func IterMap[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(k K, v V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func IterMapPairs[K comparable, V any](m map[K]V) iter.Seq[Pair[K, V]] {
	return func(yield func(kvp Pair[K, V]) bool) {
		for k, v := range m {
			if !yield(NewPair(k, v)) {
				return
			}
		}
	}
}

func IterSet[T comparable](s Set[T]) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for t := range s {
			if !yield(t) {
				return
			}
		}
	}
}

func Only[T any](v T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		yield(v)
	}
}

func Empty[T any]() iter.Seq[T] {
	return func(_ func(v T) bool) {}
}
