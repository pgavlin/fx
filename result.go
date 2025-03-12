package fx

// A Result wraps a (T, error) tuple in a single value.
type Result[T any] struct {
	v   T
	err error
}

// Unpack returns the Result's contained (T, error).
func (r Result[T]) Unpack() (T, error) {
	return r.v, r.err
}

// OK creates a Result that contains (T, nil).
func OK[T any](v T) Result[T] {
	return Result[T]{v: v}
}

// Err creates a Result that contains (_, err).
func Err[T any](e error) Result[T] {
	return Result[T]{err: e}
}

// Try creates a Result that contains (v, err).
func Try[T any](v T, e error) Result[T] {
	if e != nil {
		return Err[T](e)
	}
	return OK(v)
}

// TryFunc wraps a function that returns (U, error) so that it instead returns a Result[U].
func TryFunc[T, U any](fn func(t T) (U, error)) func(t T) Result[U] {
	return func(t T) Result[U] {
		return Try(fn(t))
	}
}
