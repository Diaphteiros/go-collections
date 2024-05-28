package iterators_test

import (
	"github.com/Diaphteiros/go-collections/pkg/iterators"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArrayIterator Tests", func() {

	It("should iterate over the slice", func() {
		example := []int{1, 2, 3}
		it := iterators.NewArrayIterator(&example, 0, len(example))
		Expect(it.HasNext()).To(BeTrue())
		Expect(it.Next()).To(Equal(1))
		Expect(it.HasNext()).To(BeTrue())
		Expect(it.Next()).To(Equal(2))
		Expect(it.HasNext()).To(BeTrue())
		Expect(it.Next()).To(Equal(3))
		Expect(it.HasNext()).To(BeFalse())
	})

	It("should return an 'empty' iterator for an empty slice", func() {
		example := []int{}
		it := iterators.NewArrayIterator(&example, 0, len(example))
		Expect(it.HasNext()).To(BeFalse())
	})

})
