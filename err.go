package fx

type Result[T any] interface {
	Value() T
	Err() error
}

type ok[T any] struct {
	v T
}

func (o ok[T]) Value() T {
	return o.v
}

func (o ok[T]) Err() error {
	return nil
}

type err[T any] struct {
	err error
}

func (e err[T]) Value() (v T) {
	return
}

func (e err[T]) Err() error {
	return e.err
}

func OK[T any](v T) Result[T] {
	return ok[T]{v: v}
}

func Err[T any](e error) Result[T] {
	return err[T]{err: e}
}
