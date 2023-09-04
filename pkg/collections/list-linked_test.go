package collections

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinkedList Integrity Validation", func() {

	It("should not violate the list's integrity", func() {
		li := NewLinkedList[int](1, 2, 3, 4)
		li.validateIntegrity("after creation")
		li.Add(5)
		li.validateIntegrity("after Add")
		li.Remove(2)
		li.validateIntegrity("after Remove")
		li.RemoveIndex(0)
		li.validateIntegrity("after RemoveIndex(0)")
		li.RemoveIndex(li.size - 1)
		li.validateIntegrity("after RemoveIndex(size-1)")
		li.AddIndex(0, 0)
		li.validateIntegrity("after AddIndex to beginning of list")
		li.AddIndex(8, li.size)
		li.validateIntegrity("after AddIndex to end of list")
	})

})

// validateIntegrity checks that for each element e
// - e.next.prev == e
// - e.prev.next == e
func (l *LinkedList[T]) validateIntegrity(msg string) {
	elem := l.dummy
	Expect(elem.next).ToNot(BeNil(), msg)
	Expect(elem.prev).ToNot(BeNil(), msg)
	Expect(elem.next.prev).To(Equal(elem), msg)
	Expect(elem.prev.next).To(Equal(elem), msg)
	elem = elem.next
	for elem != l.dummy {
		Expect(elem.next).ToNot(BeNil(), msg)
		Expect(elem.prev).ToNot(BeNil(), msg)
		Expect(elem.next.prev).To(Equal(elem), msg)
		Expect(elem.prev.next).To(Equal(elem), msg)
		elem = elem.next
	}
}
