package fx

type range_ struct {
	i, max int
	any    bool
}

func (r *range_) Value() int {
	return r.i
}

func (r *range_) Next() bool {
	if r.i >= r.max {
		return false
	}
	if !r.any {
		r.any = true
	} else {
		r.i++
	}
	return true
}

func Range(min, max int) Iterator[int] {
	return &range_{i: min, max: max}
}
