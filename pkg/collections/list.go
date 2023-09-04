package collections

import (
	"encoding/json"
	"fmt"
)

type IndexOutOfBoundsError struct {
	Index int
}

var _ error = &IndexOutOfBoundsError{}

func newIndexOutOfBoundsError(i int) *IndexOutOfBoundsError {
	return &IndexOutOfBoundsError{
		Index: i,
	}
}

func (err *IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index out of bounds: %d", err.Index)
}

type List[T any] interface {
	Collection[T]
	json.Marshaler
	json.Unmarshaler

	// AddIndex adds the given element at the specified index.
	// Returns an IndexOutOfBoundsError if the index is not within the list's size or equal to l.Size().
	AddIndex(element T, idx int) error

	// RemoveIndex removes the element at the specified index.
	// Returns an IndexOutOfBoundsError if the index is not within the list's size.
	RemoveIndex(idx int) error

	// Get returns the element at the specified index.
	// Returns an IndexOutOfBoundsError if the index is not within the list's size.
	Get(idx int) (T, error)
}
