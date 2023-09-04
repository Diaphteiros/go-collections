package collections

var _ Iterator[any] = &ArrayIterator[any]{}

type ArrayIterator[T any] struct {
	data    *[]T
	current int
	limit   int
}

// NewArrayIterator returns a new ArrayIterator.
// The first element returned will be (*data)[start], the last one will be (*data)[end-1].
func NewArrayIterator[T any](data *[]T, start, end int) *ArrayIterator[T] {
	return &ArrayIterator[T]{
		data:    data,
		current: start,
		limit:   end,
	}
}

func (i *ArrayIterator[T]) Next() T {
	res := (*i.data)[i.current]
	i.current++
	return res
}

func (i *ArrayIterator[T]) HasNext() bool {
	return i.current < i.limit
}
