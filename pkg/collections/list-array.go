package collections

import iterators "github.com/Diaphteiros/go-collections/pkg/iterators"

var _ List[any] = &ArrayList[any]{}

type ArrayList[T any] struct {
	abstractList[T]
	internal []T
	size     int
}

func NewArrayList[T any](elements ...T) *ArrayList[T] {
	res := &ArrayList[T]{
		internal: make([]T, len(elements)),
	}

	res.abstractCollection.funcIterator = res.Iterator
	res.funcAdd = res.Add
	res.funcClear = res.Clear
	res.funcRemove = res.Remove
	res.funcRemoveIf = res.RemoveIf
	res.funcSize = res.Size

	res.Add(elements...)

	return res
}

func NewArrayListFromCollection[T any](c Collection[T]) *ArrayList[T] {
	res := NewArrayList[T]()
	res.AddAll(c)
	return res
}

///////////////////////////////
// COLLECTION IMPLEMENTATION //
///////////////////////////////

// Add ensures that this collection contains the specified elements.
// Returns true if the collection changed as a result of the operation.
// (returns false if the collection does not support duplicates and already contained all given elements).
func (l *ArrayList[T]) Add(elements ...T) bool {
	if len(elements) == 0 {
		return false
	}
	if len(l.internal) < l.Size()+len(elements) {
		l.resize(maximum(len(l.internal)*2, len(elements)))
	}
	for i := range elements {
		l.internal[l.size] = elements[i]
		l.size++
	}
	return true
}

// AddAll adds all elements from the given collection to this one.
// Returns true if the collection changed as a result of the operation.
func (l *ArrayList[T]) AddAll(c Collection[T]) bool {
	it := c.Iterator()
	changed := false
	for it.HasNext() {
		changed = l.Add(it.Next()) || changed
	}
	return changed
}

// Clear removes all of the elements from this collection.
func (l *ArrayList[T]) Clear() {
	l.internal = []T{}
	l.size = 0
}

// Returns an iterator over the elements in this collection.
func (l *ArrayList[T]) Iterator() iterators.Iterator[T] {
	return iterators.NewArrayIterator[T](&l.internal, 0, l.size)
}

// Removes all given elements from the collection, if they are present.
// Each specified element is removed only once, even if it is contained multiple times in the collection.
// Returns true if the collection changed as a result of the operation.
func (l *ArrayList[T]) Remove(elements ...T) bool {
	return l.remove(false, elements...)
}

// RemoveAllOf removes all instances of the given elements from the collection.
// Returns true if the collection changed as a result of the operation.
func (l *ArrayList[T]) RemoveAllOf(elements ...T) bool {
	return l.remove(true, elements...)
}

// RemoveIf removes all elements from the collection that satisfy the given predicate.
// Returns true if the collection changed as a result of the operation.
func (l *ArrayList[T]) RemoveIf(filter Predicate[T]) bool {
	changed := false
	for i := 0; i < l.size; i++ {
		if filter(l.internal[i]) {
			l.removeIndex(i, false)
			i--
			changed = true
		}
	}
	return changed
}

// Size returns the number of elements in this collection.
func (l *ArrayList[T]) Size() int {
	return l.size
}

// ToSlice returns a slice containing the elements in this collection.
func (l *ArrayList[T]) ToSlice() []T {
	res := make([]T, l.size)
	for i := range res {
		res[i] = l.internal[i]
	}
	return res
}

/////////////////////////
// LIST IMPLEMENTATION //
/////////////////////////

func (l *ArrayList[T]) AddIndex(element T, idx int) error {
	if idx < 0 || idx > l.Size() {
		return newIndexOutOfBoundsError(idx)
	}
	if len(l.internal) <= l.Size()+1 {
		l.resize(maximum(len(l.internal), 1) * 2)
	}
	for i := l.Size() - 1; i >= idx; i-- {
		l.internal[i+1] = l.internal[i]
	}
	l.internal[idx] = element
	l.size++
	return nil
}

func (l *ArrayList[T]) RemoveIndex(idx int) error {
	if idx < 0 || idx >= l.Size() {
		return newIndexOutOfBoundsError(idx)
	}
	l.removeIndex(idx, true)
	return nil
}

func (l *ArrayList[T]) Get(idx int) (T, error) {
	var res T
	if idx < 0 || idx >= l.Size() {
		return res, newIndexOutOfBoundsError(idx)
	}
	return l.internal[idx], nil
}

/////////////////////////
// AUXILIARY FUNCTIONS //
/////////////////////////

func (l *ArrayList[T]) resize(newSize int) {
	res := make([]T, newSize)
	for i := 0; i < len(l.internal) && i < len(res); i++ {
		res[i] = l.internal[i]
	}
	l.internal = res
}

// remove removes all specified elements.
// If all is true, all instances of these elements are removed, otherwise only one.
// Resizes the list if it gets smaller than half the internal size.
func (l *ArrayList[T]) remove(all bool, elements ...T) bool {
	changed := false
	tbri := 0
	for tbri < len(elements) {
		idx := -1
		for i := 0; i < l.size; i++ {
			if Equals(l.internal[i], elements[tbri]) {
				idx = i
				break
			}
		}
		if idx >= 0 {
			l.removeIndex(idx, false)
			changed = true
		}
		if !all || idx < 0 {
			tbri++
		}
	}
	if l.size < len(l.internal)/2 {
		l.resize(len(l.internal) / 2)
	}
	return changed
}

func (l *ArrayList[T]) removeIndex(idx int, resize bool) {
	if idx < 0 || idx >= l.size {
		return
	}
	for i := idx + 1; i < l.size; i++ {
		l.internal[i-1] = l.internal[i]
	}
	l.size--
	if resize {
		if l.size < len(l.internal)/2 {
			l.resize(len(l.internal) / 2)
		}
	}
}

func maximum(elems ...int) int {
	if len(elems) == 0 {
		return 0
	}
	max := elems[0]
	for _, e := range elems[1:] {
		if e > max {
			max = e
		}
	}
	return max
}
